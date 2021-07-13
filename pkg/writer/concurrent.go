package writer

import "github.com/ArminGodiz/golang-code-challenge/pkg/models"

// ConcurrentWriter there is 1_3 goRoutines and we use them to convert struct to string
type ConcurrentWriter struct {
	GoRoutinesCapacity int
	InputChannel       chan models.CsvData
	OutputChannel      chan []string
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
