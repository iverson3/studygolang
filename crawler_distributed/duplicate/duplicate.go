package duplicate

import (
	"context"
	"github.com/go-redis/redis"
	"log"
	"regexp"
)

// redis服务实现url去重
const key = "crawler_zhenai_profile_url_set"
var idUrlRe = regexp.MustCompile(`http[s]?://album.zhenai.com/u/[\d]+`)
var ctx = context.Background()

func IsDuplicate(clientChan chan *redis.Client, url string) bool {
	// 查重只查用户url
	match := idUrlRe.FindStringSubmatch(url)
	if len(match) == 0 {
		log.Printf("not profile url: %s", url)
		return false
	}

	client := <- clientChan
	_, err := client.HGet(ctx, key, url).Result()
	if err == nil {
		log.Printf("yes yes yes yes yes url: %s", url)
		return true
	} else {
		log.Printf("xxxxxxxxxxxxxxxxxxxxxxxx url: %s", url)
		// 不存在redis中 则存入
		//err = client.HSet(key, url, "1").Err()
		//if err != nil {
		//	log.Printf("redis hash-set error: %v", err)
		//}
		return false
	}
}

func CreateRedisClientPool(host string, count int) (chan *redis.Client, int) {
	var clients []*redis.Client
	for i := 0; i < count; i++ {
		client := redis.NewClient(&redis.Options{
			Addr:     host,
			Password: "",
			DB:       0,
		})
		//defer client.Close()

		pong, err := client.Ping(ctx).Result()
		if err != nil || pong != "PONG" {
			log.Printf("Error connecting to %s: %v", host, err)
		} else {
			clients = append(clients, client)
			log.Printf("Connected to %s.", host)
		}
	}
	// 如果一个redisClient都没创建起来
	if len(clients) == 0 {
		return nil, 0
	}

	out := make(chan *redis.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out, len(clients)
}






func IsDuplicate2(client *redis.Client, url string) bool {
	// 查重只查用户url
	match := idUrlRe.FindStringSubmatch(url)
	if len(match) == 0 {
		log.Println("no profile url")
		return false
	}

	_, err := client.HGet(ctx, key, url).Result()
	if err == nil {
		log.Printf("url is exist-------------; %s", url)
		return true
	} else {
		// 不存在redis中 则存入
		err = client.HSet(ctx, key, url, "1").Err()
		if err != nil {
			log.Printf("redis hash-set error: %v", err)
		}
		return false
	}
}

func CreateRedisClient2(host string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "",
		DB:       0,
	})
	//defer client.Close()

	pong, err := client.Ping(ctx).Result()
	if err != nil || pong != "PONG" {
		log.Printf("Error connecting to %s: %v", host, err)
		return nil, err
	}
	log.Printf("Connected to %s.", host)
	return client, nil
}
