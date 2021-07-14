package broker

import (
	"github.com/ArminGodiz/golang-code-challenge/pkg/models"
	"net/http"
	"net/url"
	"strconv"
	"testing"
)

func TestBroker_Consume(t *testing.T) {
	testOutputChannel := make(chan models.BrokerData, 2)
	brokerTestObject := broker{}
	go brokerTestObject.Consume(testOutputChannel, 4040)
	testBrokerData := models.BrokerData{UserName: "Armin", ID: 1, TrafficUsage: 504, Ip: "1234", Port: "4040"}
	data := url.Values{
		"user_name":     {testBrokerData.UserName},
		"id":            {strconv.Itoa(testBrokerData.ID)},
		"traffic_usage": {strconv.Itoa(testBrokerData.TrafficUsage)},
		"ip":            {testBrokerData.Ip},
		"port":          {testBrokerData.Port},
	}
	_, err := http.PostForm("http://localhost:4040/broker", data)
	if err != nil {
		panic(err)
	}
	result := <-testOutputChannel
	if result.UserName != testBrokerData.UserName || result.ID != testBrokerData.ID || result.TrafficUsage != testBrokerData.TrafficUsage || result.Port != testBrokerData.Port || result.Ip != testBrokerData.Ip {
		t.Errorf("Error in consuming or passing data")
	}
}
