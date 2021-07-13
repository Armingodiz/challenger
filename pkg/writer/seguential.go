package writer

import (
	"github.com/ArminGodiz/golang-code-challenge/pkg/models"
	"strconv"
)

// SequentialWriter There is 0 goRoutines and we write without using any goRoutine
type SequentialWriter struct {
	InputChannel  chan models.CsvData
	OutputChannel chan []string
}

func (w *SequentialWriter) StartWriting() {
	caches := make(map[int]WritingCache)
	for input := range w.InputChannel {
		manageCaches(caches, w.OutputChannel, getTypeData(input), convertToString(input))
	}
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
