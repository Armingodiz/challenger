package cmd

import (
	"fmt"
	pkgCombiner "github.com/ArminGodiz/golang-code-challenge/pkg/combiner"
	"github.com/ArminGodiz/golang-code-challenge/pkg/models"
	pkgReader "github.com/ArminGodiz/golang-code-challenge/pkg/reader"
	pkgWriter "github.com/ArminGodiz/golang-code-challenge/pkg/writer"
)

func StartApplication() {
	//set config
	config := GetConfig()
	redisPort := config.RedisPort
	goroutinesCount := config.GoroutinesCount
	combinerRoutines, writerRoutines := setGoroutines(goroutinesCount)
	//create channels
	brokerChannel := make(chan models.BrokerData, 200)
	combinerChannel := make(chan models.CsvData, 200)
	writerChannel := make(chan []string, 200)
	//create each part
	reader := &pkgReader.Reader{}
	combiner := pkgCombiner.GetCombiner(combinerRoutines, brokerChannel, combinerChannel)
	writer := pkgWriter.GetNewWriter(writerRoutines, combinerChannel, writerChannel)
	//start each part in go routine
	go reader.StartReading(brokerChannel)
	go combiner.StartCombining(redisPort)
	go writer.StartWriting()
	//waiting in main goRoutine for response from last part(writer)
	for result := range writerChannel {
		if result != nil {
			fmt.Println("part written on file")
		}
	}
}

func setGoroutines(count int) (int, int) {
	if (count-1)%2 == 0 {
		return (count - 1) / 2, (count - 1) / 2
	} else {
		return (count)/2 - 1, count / 2
	}
}
