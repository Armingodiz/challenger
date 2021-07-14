package combiner

import (
	pkgCache "github.com/ArminGodiz/golang-code-challenge/pkg/cache"
	"github.com/ArminGodiz/golang-code-challenge/pkg/models"
)

type CombinerMock struct {
	GoRoutinesCapacity int
	InputChannel       chan models.BrokerData
	OutputChannel      chan models.CsvData
	Cache              pkgCache.CacheInterface
}

func (combiner *CombinerMock) StartCombining(port int) {
	for i := 0; i < combiner.GoRoutinesCapacity; i++ {
		go func() {
			for brokerData := range combiner.InputChannel {
				mac := ""
				combiner.OutputChannel <- models.CsvData{BrokerInfo: brokerData, Mac: mac}
			}
		}()
	}
}
