package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

// 全局的redis连接池实例
var pool *redis.Pool

// 初始化redis连接池
func initPool(address string, pwd string, maxIdle, maxActive int, idleTimeout time.Duration)  {
	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address, redis.DialPassword(pwd))
		},
		MaxIdle:         maxIdle,     // 最大空闲连接数
		MaxActive:       maxActive,   // 与redis的最大连接数，0表示无限制
		IdleTimeout:     idleTimeout, // 最大空闲时间
	}
}