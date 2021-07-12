package reader

import (
	"github.com/ArminGodiz/golang-code-challenge/pkg/broker"
	"github.com/ArminGodiz/golang-code-challenge/pkg/models"
)

type Reader struct {
	Broker broker.Broker
}

func (reader *Reader) StartReading(inputChannel chan models.BrokerData) {
	broker.BrokerObject.Consume(inputChannel)
}
