package main

import (
	"fmt"
	"time"
)

func writeNums(max int, in chan int) {
	for i := 1; i <= max; i++ {
		in <- i
	}
	close(in)
}

func readAndDeal(out chan int, res chan int, flag chan bool) {
	for num := range out {

		is := true
		for i := 2; i < num; i++ {
			if num % i == 0 {
				is = false
				break
			}
		}
		if is {
			res <- num
		}

	}
	flag<- true
	fmt.Println("一个worker线程退出")
}

// 用goroutine和channel 实现找出 1-n之间的所有质数

func main() {
	fmt.Println(time.Now())
	max := 500000
	goNum := 1200
	numsChan := make(chan int)
	resChan := make(chan int, max)
	flagChan := make(chan bool, goNum)

	go writeNums(max, numsChan)
	for i := 0; i < goNum; i++ {
		go readAndDeal(numsChan, resChan, flagChan)
	}

	//flag := 0
	//for {
	//	<-flagChan
	//	flag++
	//	if flag == goNum {
	//		break
	//	}
	//}

	// 开个goroutine来处理
	go func() {
		for i := 0; i < goNum; i++ {
			<-flagChan
		}
		close(resChan)
	}()

	fmt.Println("result:")
	for {
		v, ok := <-resChan
		if !ok {
			break
		}
		fmt.Println(v)
	}
	//for v := range resChan {
	//	fmt.Println(v)
	//}

	fmt.Println(time.Now())
}
