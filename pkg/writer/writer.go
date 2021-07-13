package writer

import (
	"encoding/csv"
	"github.com/ArminGodiz/golang-code-challenge/pkg/models"
	"log"
	"os"
	"strconv"
)

type WriterInterface interface {
	StartWriting()
}

type MiddleWare struct {
	Data string
	Type int
}
type WritingCache []string

// SequentialWriter There is 0 goRoutines and we write without using any goRoutine
type SequentialWriter struct {
	InputChannel  chan models.CsvData
	OutputChannel chan []string
}

// ConcurrentWriter there is 1_3 goRoutines and we use them to convert struct to string
type ConcurrentWriter struct {
	GoRoutinesCapacity int
	InputChannel       chan models.CsvData
	OutputChannel      chan []string
}

// MultiGoroutinesWriter count of  goRoutines ==4, we set a goRoutine for each file(converting and writing will be done in GoRoutine)
type MultiGoroutinesWriter struct {
	InputChannel  chan models.CsvData
	OutputChannel chan []string
}

// HighConcurrentWriter count of  goRoutines >=8 , we set a goRoutine for each file(just for writing) and other goRoutines will be used as worker pool for
// converting struct to string
type HighConcurrentWriter struct {
	GoRoutinesCapacity int
	InputChannel       chan models.CsvData
	OutputChannel      chan []string
}

func GetNewWriter(goRoutinesCapacity int, inputChannel chan models.CsvData, outputChannel chan []string) WriterInterface {
	if goRoutinesCapacity == 0 {
		return &SequentialWriter{
			InputChannel:  inputChannel,
			OutputChannel: outputChannel,
		}
	} else if goRoutinesCapacity > 0 && goRoutinesCapacity < 4 {
		return &ConcurrentWriter{
			GoRoutinesCapacity: goRoutinesCapacity,
			InputChannel:       inputChannel,
			OutputChannel:      outputChannel,
		}
	} else if goRoutinesCapacity == 4 {
		return &MultiGoroutinesWriter{
			InputChannel:  inputChannel,
			OutputChannel: outputChannel,
		}
	} else if goRoutinesCapacity >= 8 {
		return &HighConcurrentWriter{
			GoRoutinesCapacity: goRoutinesCapacity,
			InputChannel:       inputChannel,
			OutputChannel:      outputChannel,
		}
	} else {
		return nil
	}
}
func (w *SequentialWriter) StartWriting() {
	caches := make(map[int]WritingCache)
	for input := range w.InputChannel {
		manageCaches(caches, w.OutputChannel, getTypeData(input), convertToString(input))
	}
}

func (w *ConcurrentWriter) StartWriting() {
	middleWare := make(chan MiddleWare, 200)
	for i := 0; i < w.GoRoutinesCapacity; i++ {
		go func() {
			for data := range w.InputChannel {
				middleWare <- MiddleWare{Data: convertToString(data), Type: getTypeData(data)}
			}
		}()
	}
	caches := make(map[int]WritingCache)
	for converted := range middleWare {
		manageCaches(caches, w.OutputChannel, converted.Type, converted.Data)
	}
}

func (w *MultiGoroutinesWriter) StartWriting() {
}

func (w *HighConcurrentWriter) StartWriting() {
}

func convertToString(data models.CsvData) string {
	return data.BrokerInfo.UserName + "|" + strconv.Itoa(data.BrokerInfo.ID) + "|" + strconv.Itoa(data.BrokerInfo.TrafficUsage) + "|" + data.BrokerInfo.Ip + "|" + data.BrokerInfo.Port + "|" + data.Mac
}
func manageCaches(caches map[int]WritingCache, out chan []string, dataType int, data string) {
	caches[dataType-1] = append(caches[dataType-1], data)
	for i := 0; i < 4; i++ {
		if len(caches[i]) >= 5 {
			WriteToFile(caches[i], getPath(i+1))
			out <- caches[i]
			caches[i] = *new(WritingCache)
		}
	}
}

func getTypeData(data models.CsvData) int {
	if data.BrokerInfo.TrafficUsage >= 0 && data.BrokerInfo.TrafficUsage <= 100 {
		return 1
	} else if data.BrokerInfo.TrafficUsage >= 101 && data.BrokerInfo.TrafficUsage <= 500 {
		return 2
	} else if data.BrokerInfo.TrafficUsage >= 501 && data.BrokerInfo.TrafficUsage <= 1000 {
		return 3
	} else if data.BrokerInfo.TrafficUsage >= 1001 && data.BrokerInfo.TrafficUsage <= 1500 {
		return 4
	} else {
		return 0
	}
}
func getPath(dataType int) string {
	switch dataType {
	case 1:
		return "0_100.csv"
	case 2:
		return "101_500.csv"
	case 3:
		return "501_1000.csv"
	case 4:
		return "1001_1500.csv"
	default:
		return ""
	}
}

func WriteToFile(cache WritingCache, path string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	defer f.Close()
	if err != nil {
		log.Fatalln("failed to open or create file", err)
	}
	w := csv.NewWriter(f)
	for _, record := range cache {
		data := make([]string, 1)
		data[0] = record
		if err := w.Write(data); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
	w.Flush()
	err = w.Error()
	if err != nil {
		panic(err)
	}
	err = f.Sync()
	if err != nil {
		panic(err)
	}
}
