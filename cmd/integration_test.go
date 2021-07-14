package cmd

import (
	pkgCombiner "github.com/ArminGodiz/golang-code-challenge/pkg/combiner"
	"github.com/ArminGodiz/golang-code-challenge/pkg/models"
	pkgReader "github.com/ArminGodiz/golang-code-challenge/pkg/reader"
	pkgWriter "github.com/ArminGodiz/golang-code-challenge/pkg/writer"
	"testing"
)

func TestIntegration(t *testing.T) {
	brokerChannel := make(chan models.BrokerData, 5)
	combinerChannel := make(chan models.CsvData, 5)
	writerChannel := make(chan []string, 1)
	//create each part
	reader := &pkgReader.Reader{}
	combiner := &pkgCombiner.CombinerMock{2, brokerChannel, combinerChannel, nil}
	writer := pkgWriter.GetNewWriter(2, combinerChannel, writerChannel)
	//start each part in go routine
	go reader.StartReading(brokerChannel, 8080)
	go combiner.StartCombining(2020)
	go writer.StartWriting()
	testConnection(brokerChannel, writerChannel)
}
func testConnection(in chan models.BrokerData, out chan []string) bool {
	testData := []models.BrokerData{
		models.BrokerData{UserName: "1", ID: 1, TrafficUsage: 1, Ip: "2020", Port: "8080"},
		models.BrokerData{UserName: "2", ID: 1, TrafficUsage: 1, Ip: "2020", Port: "8080"},
		models.BrokerData{UserName: "3", ID: 1, TrafficUsage: 1, Ip: "2020", Port: "8080"},
		models.BrokerData{UserName: "4", ID: 1, TrafficUsage: 1, Ip: "2020", Port: "8080"},
		models.BrokerData{UserName: "5", ID: 1, TrafficUsage: 1, Ip: "2020", Port: "8080"},
	}
	for i := 0; i < 5; i++ {
		in <- testData[i]
	}
	res := <-out
	if len(res) == 0 {
		return false
	}
	return true
}
