package combiner

import (
	pkgCache "github.com/ArminGodiz/golang-code-challenge/pkg/cache"
	"github.com/ArminGodiz/golang-code-challenge/pkg/models"
	"testing"
)

func TestGetCombiner(t *testing.T) {
	testInpChan := make(chan models.BrokerData)
	testOutChan := make(chan models.CsvData)
	testCombiner := GetCombiner(1, testInpChan, testOutChan)
	if testCombiner.InputChannel != testInpChan || testCombiner.OutputChannel != testOutChan {
		t.Errorf("Error in Creating Combiner")
	}
}

func TestCombiner_StartCombining(t *testing.T) {
	testInpChan := make(chan models.BrokerData, 1)
	testOutChan := make(chan models.CsvData, 1)
	testCombiner := CombinerMock{3, testInpChan, testOutChan, pkgCache.GetMockCache()}
	testCombiner.StartCombining(3030)
	testData := models.BrokerData{UserName: "Armin", ID: 1, TrafficUsage: 504, Ip: "1234", Port: "4040"}
	testInpChan <- testData
	result := <-testOutChan
	if result.BrokerInfo != testData {
		t.Errorf("Error in combining data")
	}
}
