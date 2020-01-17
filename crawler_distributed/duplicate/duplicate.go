package duplicate

import (
	"github.com/go-redis/redis"
	"log"
	"regexp"
)

// redis服务实现url去重
const key = "crawler_zhenai_profile_url_set"
var idUrlRe = regexp.MustCompile(`http[s]?://album.zhenai.com/u/[\d]+`)

func IsDuplicate(client *redis.Client, url string) bool {
	// 查重只查用户url
	match := idUrlRe.FindStringSubmatch(url)
	if len(match) == 0 {
		log.Println("no profile url")
		return false
	}

	_, err := client.HGet(key, url).Result()
	if err == nil {
		log.Printf("url is exist-------------; %s", url)
		return true
	} else {
		// 不存在redis中 则存入
		err = client.HSet(key, url, "1").Err()
		if err != nil {
			log.Printf("redis hash-set error: %v", err)
		}
		return false
	}
}

func CreateRedisClient(host string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "",
		DB:       0,
	})
	//defer client.Close()

	pong, err := client.Ping().Result()
	if err != nil || pong != "PONG" {
		return nil, err
	}
	return client, nil
}