package writer

import (
	"encoding/csv"
	"fmt"
	"github.com/ArminGodiz/golang-code-challenge/pkg/models"
	"log"
	"os"
)

type WriterInterface interface {
	StartWriting()
}

type MiddleWare struct {
	Data string
	Type int
}
type WritingCache []string

func GetNewWriter(goRoutinesCapacity int, inputChannel chan models.CsvData, outputChannel chan []string) WriterInterface {
	if goRoutinesCapacity == 0 {
		return &SequentialWriter{
			InputChannel:  inputChannel,
			OutputChannel: outputChannel,
		}
	} else if goRoutinesCapacity > 0 && goRoutinesCapacity < 5 {
		return &ConcurrentWriter{
			GoRoutinesCapacity: goRoutinesCapacity,
			InputChannel:       inputChannel,
			OutputChannel:      outputChannel,
		}
	} else if goRoutinesCapacity == 5 {
		return &MultiGoroutinesWriter{
			InputChannel:  inputChannel,
			OutputChannel: outputChannel,
		}
	} else if goRoutinesCapacity >= 9 {
		return &HighConcurrentWriter{
			GoRoutinesCapacity: goRoutinesCapacity,
			InputChannel:       inputChannel,
			OutputChannel:      outputChannel,
		}
	} else {
		return nil
	}
}

func WriteToFile(cache WritingCache, path string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	defer f.Close()
	if err != nil {
		//log.Fatalln("failed to open or create file", err)
		fmt.Println(err)
	} else {
		w := csv.NewWriter(f)
		for _, record := range cache {
			data := make([]string, 1)
			data[0] = record
			if err := w.Write(data); err != nil {
				log.Fatalln("error writing record to file", err)
				//fmt.Println(err)
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
		return "output/0_100.csv"
	case 2:
		return "output/101_500.csv"
	case 3:
		return "output/501_1000.csv"
	case 4:
		return "output/1001_1500.csv"
	default:
		return ""
	}
}
