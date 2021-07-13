package cmd

import (
	"fmt"
	pkgCombiner "github.com/ArminGodiz/golang-code-challenge/pkg/combiner"
	"github.com/ArminGodiz/golang-code-challenge/pkg/models"
	pkgReader "github.com/ArminGodiz/golang-code-challenge/pkg/reader"
	pkgWriter "github.com/ArminGodiz/golang-code-challenge/pkg/writer"
)

func StartApplication() {
	redisPort := 8282
	brokerChannel := make(chan models.BrokerData, 200)
	combinerChannel := make(chan models.CsvData, 200)
	writerChannel := make(chan []string, 200)
	reader := &pkgReader.Reader{}
	combiner := pkgCombiner.GetCombiner(4, brokerChannel, combinerChannel)
	writer := pkgWriter.GetNewWriter(9, combinerChannel, writerChannel)
	go reader.StartReading(brokerChannel)
	go combiner.StartCombining(redisPort)
	go writer.StartWriting()
	for result := range writerChannel {
		if result != nil {
			fmt.Println("part written on file")
		}
	}
}
