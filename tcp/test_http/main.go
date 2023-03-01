package main

import (
	"context"
	"fmt"
	"studygolang/tcp/test_http/util"
	"sync"
)

func sum(n int) int {
	if n == 1 {
		return 1
	}
	return n + sum(n-1)
}

func main() {
	// 求 1+2+...+100 的结果
	res := sum(100)
	fmt.Println(res)


	// 求 1+2+...+100 的结果
	var res2 int
	for i := 1; i <= 100; i++ {
		res2 += i
	}
	fmt.Println(res2)

	return
	var wg sync.WaitGroup
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go job(&wg, i+1)
	}

	wg.Wait()
	fmt.Println("job all over")
}

func job(wg *sync.WaitGroup, n int) {
	fmt.Printf("prepare to request: %d\n", n)

	url := fmt.Sprintf("http://81.69.56.251:9099/go-yx-extension/api/v1/test/test/getname?name=%d", n)
	respBytes, err := util.NewHttpClient(nil, 10).GET(context.Background(), url, nil)
	if err != nil {
		fmt.Printf("Get request failed, number: %d, error: %v\n", n, err)
	} else {
		fmt.Printf("Get request succeed, number: %d, resp: %s\n", n, respBytes)
	}
	wg.Done()
}