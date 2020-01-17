package duplicate

import (
	"github.com/go-redis/redis"
	"log"
	"studygolang/crawler_distributed/config"
)

func Redis() {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisServerUrl,
		Password: "",
		DB:       0,
	})
	defer client.Close()

	pong, err := client.Ping().Result()
	if err != nil || pong != "PONG" {
		panic(err)
	}

	result, err := client.HGet("crawler_zhenai_profile_url_set", "http://www.zhenai.com/zhenghun/shanghai").Result()
	if err == nil {
		log.Println(result)
	} else {
		log.Printf("error: %v", err)
	}

}