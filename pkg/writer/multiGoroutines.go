package writer

import "github.com/ArminGodiz/golang-code-challenge/pkg/models"

// MultiGoroutinesWriter count of  goRoutines ==4, we set a goRoutine for each file(converting and writing will be done in GoRoutine)
type MultiGoroutinesWriter struct {
	InputChannel  chan models.CsvData
	OutputChannel chan []string
}

func (w *MultiGoroutinesWriter) StartWriting() {
	caches := make(map[int]WritingCache)
	var channels = []chan models.CsvData{
		make(chan models.CsvData, 200),
		make(chan models.CsvData, 200),
		make(chan models.CsvData, 200),
		make(chan models.CsvData, 200),
	}
	for i := 0; i < 4; i++ {
		go routineWorker(i+1, channels[i], caches[i], w.OutputChannel)
	}
	for input := range w.InputChannel {
		dataType := getTypeData(input)
		channels[dataType-1] <- input
	}
}
func routineWorker(number int, inp chan models.CsvData, cache WritingCache, out chan []string) {
	for input := range inp {
		//fmt.Println("fadssdfaadfsfadsadsf")
		cache = append(cache, convertToString(input))
		if len(cache) >= 5 {
			WriteToFile(cache, getPath(number))
			out <- cache
			cache = *new(WritingCache)
		}
	}
}