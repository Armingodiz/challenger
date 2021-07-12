package broker

import (
	"github.com/ArminGodiz/golang-code-challenge/pkg/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Broker interface {
	Consume(chan models.BrokerData)
}
type broker struct {
	OutputChannel chan models.BrokerData
}

var (
	BrokerObject broker
)

func (b broker) Consume(outputChannel chan models.BrokerData) {
	BrokerObject.OutputChannel = outputChannel
	r := mux.NewRouter()
	r.HandleFunc("/broker", consume)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func consume(w http.ResponseWriter, r *http.Request) {
	userName := r.FormValue("user_name")
	idString := r.FormValue("id")
	id, _ := strconv.Atoi(idString)
	trafficUsageString := r.FormValue("traffic_usage")
	trafficUsage, _ := strconv.Atoi(trafficUsageString)
	ip := r.FormValue("ip")
	port := r.FormValue("port")
	newInfo := models.BrokerData{
		UserName:     userName,
		ID:           id,
		TrafficUsage: trafficUsage,
		Ip:           ip,
		Port:         port,
	}
	r.Write(w)
	//fmt.Println(newInfo)
	BrokerObject.OutputChannel <- newInfo
	//fmt.Println("sent to channel")
}
