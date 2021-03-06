package combiner

import (
	pkgCache "github.com/ArminGodiz/golang-code-challenge/pkg/cache"
	"github.com/ArminGodiz/golang-code-challenge/pkg/models"
)

type CombinerInterface interface {
	StartCombining(port int)
}

type Combiner struct {
	GoRoutinesCapacity int
	InputChannel       chan models.BrokerData
	OutputChannel      chan models.CsvData
	Cache              pkgCache.CacheInterface
}


func GetCombiner(goRoutinesCap int, inputChannel chan models.BrokerData, outputChannel chan models.CsvData) *Combiner {
	return &Combiner{
		GoRoutinesCapacity: goRoutinesCap,
		InputChannel:       inputChannel,
		OutputChannel:      outputChannel,
	}
}

func (combiner *Combiner) StartCombining(port int) {
	pkgCache.SetCacheClient(port)
	combiner.Cache = pkgCache.Object
	for i := 0; i < combiner.GoRoutinesCapacity; i++ {
		go func() {
			for brokerData := range combiner.InputChannel {
				//time.Sleep(3 * time.Second)
				mac := combiner.Cache.Get(brokerData.Ip)
				combiner.OutputChannel <- models.CsvData{BrokerInfo: brokerData, Mac: mac}
			}
		}()
	}
}


