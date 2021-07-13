package cache

import (
	"context"
	"github.com/ArminGodiz/golang-code-challenge/pkg/models"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
	"time"
)

type CacheInterface interface {
	Get(string) string
	Set(models.CacheData) error
	InitializeCache(port int) error
}

type cacheClient struct {
	Client *redis.Client
}

var Object CacheInterface

func SetCacheClient(port int) {
	addr := "localhost:" + strconv.Itoa(port)
	redisClient := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln(err)
	}
	Object = &cacheClient{Client: redisClient}
	err2 := Object.InitializeCache(port)
	if err2 != nil {
		panic(err2)
	}
}

func (c *cacheClient) Get(key string) string {
	value, err := c.Client.Get(context.Background(), key).Result()
	if err != nil {
		panic(err)
	}
	return value
}

func (c *cacheClient) Set(ipMac models.CacheData) error {
	err := c.Client.Set(context.Background(), ipMac.Ip, ipMac.Mac, time.Hour).Err()
	return err
}
