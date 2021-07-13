package writer

import "github.com/ArminGodiz/golang-code-challenge/pkg/models"

// HighConcurrentWriter count of  goRoutines >=8 , we set a goRoutine for each file(just for writing) and other goRoutines will be used as worker pool for
// converting struct to string
type HighConcurrentWriter struct {
	GoRoutinesCapacity int
	InputChannel       chan models.CsvData
	OutputChannel      chan []string
}

func (w *HighConcurrentWriter) StartWriting() {
	var middleWares = []chan MiddleWare{
		make(chan MiddleWare, 200),
		make(chan MiddleWare, 200),
		make(chan MiddleWare, 200),
		make(chan MiddleWare, 200),
	}
	for i := 0; i < w.GoRoutinesCapacity-4; i++ {
		go func() {
			for data := range w.InputChannel {
				dataType := getTypeData(data)
				middleWares[dataType-1] <- MiddleWare{Data: convertToString(data), Type: dataType}
			}
		}()
	}
	caches := make(map[int]WritingCache)
	for i := 0; i < 4; i++ {
		go writerWorker(middleWares[i], caches[i], w.OutputChannel)
	}
}
func writerWorker(inp chan MiddleWare, cache WritingCache, out chan []string) {
	for input := range inp {
		cache = append(cache, input.Data)
		if len(cache) >= 5 {
			WriteToFile(cache, getPath(input.Type))
			out <- cache
			cache = *new(WritingCache)
		}
	}
}
