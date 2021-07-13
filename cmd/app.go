package cmd

import (
	"fmt"
	pkgCombiner "github.com/ArminGodiz/golang-code-challenge/pkg/combiner"
	"github.com/ArminGodiz/golang-code-challenge/pkg/models"
	pkgReader "github.com/ArminGodiz/golang-code-challenge/pkg/reader"
	pkgWriter "github.com/ArminGodiz/golang-code-challenge/pkg/writer"
)

func StartApplication() {
	brokerChannel := make(chan models.BrokerData, 200)
	combinerChannel := make(chan models.CsvData, 200)
	writerChannel := make(chan models.CsvData, 200)
	reader := &pkgReader.Reader{}
	combiner := pkgCombiner.GetCombiner(4, brokerChannel, combinerChannel)
	writer := pkgWriter.GetNewWriter(2, combinerChannel, writerChannel)
	go reader.StartReading(brokerChannel)
	go combiner.StartCombining()
	go writer.StartWriting()
	for result := range writerChannel {
		fmt.Println(result)
	}
}
