package combiner

import (
	"github.com/ArminGodiz/golang-code-challenge/pkg/models"
	"testing"
)

func TestGetCombiner(t *testing.T) {
	testInpChan := make(chan models.BrokerData)
	testOutChan := make(chan models.CsvData)
	testCombiner := GetCombiner(1, testInpChan, testOutChan)
	if testCombiner.InputChannel !=testInpChan || testCombiner.OutputChannel!=testOutChan{
		t.Errorf("Error in Creating Combiner")
	}
}

func TestCombiner_StartCombining(t *testing.T) {
	testInpChan := make(chan models.BrokerData,1)
	testOutChan := make(chan models.CsvData,1)
	testCombiner := GetCombiner(1, testInpChan, testOutChan)
	testCombiner.StartCombining(8282)
}