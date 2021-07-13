package writer

import (
	"encoding/csv"
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
