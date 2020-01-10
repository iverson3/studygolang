package main

import (
	"fmt"
)

// goroutine之间的双向通道就是channel
// go slogan: 不要通过共享内存来通信, 要通过通信来共享内存

// channel是一等公民 可以作为参数或返回值

// chan<- int   该channel只能发数据 (向channel里面发数据)
// <-chan int   该channel只能收数据 (从channel里面收数据)

func doWork(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
		// 单独开goroutine来发数据到done中 以防堵塞
		go func() { done <- true }()
	}
}

type worker struct {
	in chan int
	done chan bool
}

func createWorker(id int) worker {
	w := worker{
		in: make(chan int),
		done: make(chan bool),
	}

	go doWork(id, w.in, w.done)
	return w
}

func channelDemo()  {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		//channels[i] = make(chan int)
		//go worker(i + 1, channels[i])

		// 另一种实现方式
		workers[i] = createWorker(i + 1)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	// wait for all of them
	for _, worker := range workers {
		<-worker.done
		<-worker.done
	}
}

func main() {
	channelDemo()
}

























