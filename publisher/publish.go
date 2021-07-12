package main

import (
	"encoding/json"
	"fmt"
	"github.com/ArminGodiz/golang-code-challenge/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

type Users struct {
	Users []models.BrokerData `json:"users"`
}

func main() {
	dataSet := getData()
	for i := 0; i < len(dataSet.Users); i++ {
		time.Sleep(2 * time.Second)
		data := url.Values{
			"user_name":     {dataSet.Users[i].UserName},
			"id":            {strconv.Itoa(dataSet.Users[i].ID)},
			"traffic_usage": {strconv.Itoa(dataSet.Users[i].TrafficUsage)},
			"ip":            {dataSet.Users[i].Ip},
			"port":          {dataSet.Users[i].Port},
		}
		_, err := http.PostForm("http://localhost:8080/broker", data)
		fmt.Print(data);fmt.Println("    == > sent")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getData() Users {
	currWd, _ := os.Getwd()
	jsonFile, err := ioutil.ReadFile(currWd + "/publisher/broker.json")
	if err != nil {
		fmt.Println(err)
	}
	var info Users
	_ = json.Unmarshal([]byte(jsonFile), &info)
	return info
}
