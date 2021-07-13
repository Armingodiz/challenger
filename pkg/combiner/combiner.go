package combiner

import (
	"fmt"
	pkgCache "github.com/ArminGodiz/golang-code-challenge/pkg/cache"
	"github.com/ArminGodiz/golang-code-challenge/pkg/models"
)

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
				fmt.Print("#")
				fmt.Print(brokerData.Ip)
				fmt.Println("#")
				mac := combiner.Cache.Get(brokerData.Ip)
				combiner.OutputChannel <- models.CsvData{BrokerInfo: brokerData, Mac: mac}
			}
		}()
	}
}
