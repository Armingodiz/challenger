package cmd

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	GoroutinesCount int `json:"goroutines_count"`
	RedisPort       int `json:"redis_port"`
	BrokerPort      int `json:"broker_port"`
	DataCount       int `json:"data_count"`
}

func GetConfig() Config {
	jsonFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	var config Config
	err = json.Unmarshal([]byte(jsonFile), &config)
	if err != nil {
		panic(err)
	}
	return config
}
