package main

import (
	"fmt"
	"math"
	"studygolang/wangdis/redis/server"
	"studygolang/wangdis/tcp"
)

func main() {
	var shardCount int = 27

	capacity := computeCapacity(shardCount)

	fmt.Println(shardCount, capacity)

	return
	tcp.ListenAndServe(&tcp.Config{Address: ":9000"}, server.MakeHandler())
}

func computeCapacity(param int) (size int) {
	if param <= 16 {
		return 16
	}
	n := param - 1
	fmt.Println(n)
	n |= n >> 1
	fmt.Println(n)
	n |= n >> 2
	fmt.Println(n)
	n |= n >> 4
	fmt.Println(n)
	n |= n >> 8
	fmt.Println(n)
	n |= n >> 16
	fmt.Println(n)
	if n < 0 {
		return math.MaxInt32
	}
	return n + 1
}