package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	b []byte
	n []int64
)

var (
	ballast2 = flag.Bool("ballast", false, "run program with Ballast")
)

func main() {
	flag.Parse()
	withBallast := *ballast2

	fmt.Println(withBallast)
	if withBallast {
		b = make([]byte, 2*1024*1024*1024)
	}

	start := time.Now()

	for i := 0; i < 4; i++ {
		for j := 0; j < 1024; j++ {
			n = make([]int64, 128 * 1024)  // 128 * 1024 * 8字节 = 1024*1024字节 = 1MB  / 2  =  0.5MB
			n[0] = int64(j)
			n[1] = int64(j)
			n[2] = int64(j)
		}
	}

	fmt.Printf("took %.02f s\n", float64(time.Since(start).Milliseconds())/1000)
}
