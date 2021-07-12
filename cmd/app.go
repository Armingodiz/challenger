package cmd

import (
	"fmt"
	pkgCombiner "github.com/ArminGodiz/golang-code-challenge/pkg/combiner"
	"github.com/ArminGodiz/golang-code-challenge/pkg/models"
	pkgReader "github.com/ArminGodiz/golang-code-challenge/pkg/reader"
)

func StartApplication() {
	brokerChannel := make(chan models.BrokerData, 200)
	combinerChannel := make(chan models.CsvData, 200)
	reader := &pkgReader.Reader{}
	combiner := pkgCombiner.GetCombiner(4, brokerChannel, combinerChannel)
	go reader.StartReading(brokerChannel)
	go combiner.StartCombining()
	for combined := range combinerChannel {
		fmt.Println(combined.Mac)
	}
}
