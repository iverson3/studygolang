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

	//_, err = conn.Do("hset", "user01", "name", "tomtom")
	//if err != nil {
	//	fmt.Println("hset failed! error: ", err)
	//	return
	//}
	//_, err = conn.Do("hset", "user01", "age", 26)
	//if err != nil {
	//	fmt.Println("hset failed! error: ", err)
	//	return
	//}
	//_, err = conn.Do("hset", "user01", "score", 98.5)
	//if err != nil {
	//	fmt.Println("hset failed! error: ", err)
	//	return
	//}

	//res, err := redis.Float64(conn.Do("hget", "user01", "score"))
	//if err != nil {
	//	fmt.Println("hget failed! error: ", err)
	//	return
	//}
	//fmt.Printf("hget res: %v \n", res)
	//
	//resMap, err := redis.StringMap(conn.Do("hgetall", "user01"))
	//if err != nil {
	//	fmt.Println("hgetall failed! error: ", err)
	//	return
	//}
	//fmt.Printf("hgetall res: %v \n", resMap)


	_, err = conn.Do("hmset", "user02", "name", "jack", "age", 18, "score", 88.2)
	if err != nil {
		fmt.Println("hmset failed! error: ", err)
		return
	}

	strings, err := redis.Strings(conn.Do("hmget", "user02", "name", "score"))
	if err != nil {
		fmt.Println("hmget failed! error: ", err)
		return
	}
	fmt.Printf("hmget res: %v \n", strings)



	// redis连接池技术
	pool := &redis.Pool{
		// 初始化redis连接
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379", redis.DialPassword("13396095889"))
		},
		MaxIdle:         8,     // 最大空闲连接数
		MaxActive:       0,     // 与redis的最大连接数，0表示无限制
		IdleTimeout:     100,   // 最大空闲时间
		Wait:            false,
		MaxConnLifetime: 0,
	}

	redisConn := pool.Get()  // 从连接池中获取一个连接
	_, err = redisConn.Do("get", "name")
	// 使用完之后，将当前连接放回连接池 (这里的Close()不是关闭连接，而是将连接放回连接池；放回后无法再次使用该连接，必须再次调用Get()重新从连接池中获取连接)
	defer redisConn.Close()

	// pool.Close()  // 关闭连接池，所有的redis连接资源全部关闭释放

}
