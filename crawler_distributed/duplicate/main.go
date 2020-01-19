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

	hLen := client.HLen("crawler_zhenai_profile_url_set")
	log.Printf("len: %v", hLen)

	result, err := client.HGet("crawler_zhenai_profile_url_set", "http://album.zhenai.com/u/1173210456").Result()
	if err == nil {
		log.Println(result)
	} else {
		log.Printf("error: %v", err)
	}

}