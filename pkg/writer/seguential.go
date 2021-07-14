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


