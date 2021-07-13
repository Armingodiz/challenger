package cache

import (
	"encoding/json"
	"fmt"
	"github.com/ArminGodiz/golang-code-challenge/pkg/models"
	"io/ioutil"
	"os"
)

var macMap map[string]string

type Macs struct {
	Macs []models.CacheData `json:"macs"`
}

func (c *cacheClient) InitializeCache(port int) error {
	macs := getData().Macs
	macMap = make(map[string]string)
	//fmt.Println(macs)
	for i := 0; i < len(macs); i++ {
		macMap[macs[i].Ip] = macs[i].Mac
		err := c.Set(macs[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func getData() Macs {
	currWd, _ := os.Getwd()
	jsonFile, err := ioutil.ReadFile(currWd + "/pkg/cache/cache.json")
	if err != nil {
		fmt.Println(err)
	}
	var info Macs
	_ = json.Unmarshal([]byte(jsonFile), &info)
	return info
}
