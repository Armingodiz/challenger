package writer

import (
	"github.com/ArminGodiz/golang-code-challenge/pkg/models"
	"testing"
)

func TestGetNewWriter(t *testing.T) {
	testInChan := make(chan models.CsvData)
	testOutChan := make(chan []string)
	testWriter := GetNewWriter(0, testInChan, testOutChan)
	if testWriter == nil {
		t.Errorf("Error in creating writer")
	}
}

func TestSequentialWriter_StartWriting(t *testing.T) {
	testInChan := make(chan models.CsvData, 5)
	testOutChan := make(chan []string, 1)
	testWriter := GetNewWriter(0, testInChan, testOutChan)
	go testWriter.StartWriting()
	if !testWrite(testInChan, testOutChan) {
		t.Errorf("Error in Sequential Writer")
	}
}
func TestConcurrentWriter_StartWriting(t *testing.T) {
	testInChan := make(chan models.CsvData, 5)
	testOutChan := make(chan []string, 1)
	testWriter := GetNewWriter(1, testInChan, testOutChan)
	go testWriter.StartWriting()
	if !testWrite(testInChan, testOutChan) {
		t.Errorf("Error in Concurrent Writer")
	}
}
func TestMultiGoroutinesWriter_StartWriting(t *testing.T) {
	testInChan := make(chan models.CsvData, 5)
	testOutChan := make(chan []string, 1)
	testWriter := GetNewWriter(4, testInChan, testOutChan)
	go testWriter.StartWriting()
	if !testWrite(testInChan, testOutChan) {
		t.Errorf("Error in MultiGoroutines Writer")
	}
}
func TestHighConcurrentWriter_StartWriting(t *testing.T) {
	testInChan := make(chan models.CsvData, 5)
	testOutChan := make(chan []string, 1)
	testWriter := GetNewWriter(9, testInChan, testOutChan)
	go testWriter.StartWriting()
	if !testWrite(testInChan, testOutChan) {
		t.Errorf("Error in High Concurrent Writer")
	}
}

func testWrite(in chan models.CsvData, out chan []string) bool {
	testData := []models.CsvData{
		{models.BrokerData{UserName: "1", ID: 1, TrafficUsage: 1, Ip: "2020", Port: "8080"}, "1"},
		{models.BrokerData{UserName: "2", ID: 1, TrafficUsage: 1, Ip: "2020", Port: "8080"}, "2"},
		{models.BrokerData{UserName: "3", ID: 1, TrafficUsage: 1, Ip: "2020", Port: "8080"}, "3"},
		{models.BrokerData{UserName: "4", ID: 1, TrafficUsage: 1, Ip: "2020", Port: "8080"}, "4"},
		{models.BrokerData{UserName: "5", ID: 1, TrafficUsage: 1, Ip: "2020", Port: "8080"}, "5"},
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
