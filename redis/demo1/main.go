package main

import (
	"context"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var ctx = context.Background()

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialDatabase(0), redis.DialPassword("13396095889"))
	if err != nil {
		fmt.Println("connect failed! error: ", err)
		return
	}
	defer conn.Close()

	_, err = conn.Do("set", "name", "jake111设置")
	if err != nil {
		fmt.Println("do something failed! error: ", err)
		return
	}

	res, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("do something failed! error: ", err)
		return
	}

	fmt.Printf("do res: %s", res)
}
