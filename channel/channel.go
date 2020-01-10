package main

import (
	"fmt"
	"time"
)

// goroutine之间的双向通道就是channel
// go slogan: 不要通过共享内存来通信, 要通过通信来共享内存

// channel是一等公民 可以作为参数或返回值

// chan<- int   该channel只能发数据 (向channel里面发数据)
// <-chan int   该channel只能收数据 (从channel里面收数据)

func worker(id int, c chan int) {
	for {
		// 从channel中收数据
		n := <-c
		fmt.Printf("Worker %d received %c\n", id, n)
	}
}

func createWorker(id int) chan int {
	c := make(chan int)
	go func() {
		for {
			// 从channel中收数据
			n := <-c
			fmt.Printf("Worker %d received %c\n", id, n)
		}
	}()
	return c
}

func channelDemo()  {
	// 定义一个channel
	//var c chan int   // c == nil

	// 创建channel
	//c := make(chan int)

	var channels [10]chan int
	for i := 0; i < 10; i++ {
		//channels[i] = make(chan int)
		//go worker(i + 1, channels[i])

		// 另一种实现方式
		channels[i] = createWorker(i + 1)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	// 向channel里面发数据
	//c <- 1
	//c <- 2

	time.Sleep(time.Millisecond * 50)
}


func worker2(id int, c chan int) {
	for {
		// 从channel中收数据 并判断是否不再发数据 ( close(c) 表示不再发数据 )
		n, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("Worker %d received %d\n", id, n)
	}

	// 另一种收数据并判断是否结束的方式 (结束则退出循环)
	//for n := range c {
	//	fmt.Printf("Worker %d received %d\n", id, n)
	//}
}

func bufferedChannel()  {
	// 给定大小为3 的缓冲区
	c := make(chan int, 3)

	go worker2(1, c)

	c <- 1
	c <- 2
	c <- 3
	c <- 4
	time.Sleep(time.Millisecond * 10)
}

func channelClose()  {
	c := make(chan int)
	go worker2(1, c)
	c <- 1
	c <- 2
	c <- 3
	close(c)
	time.Sleep(time.Millisecond * 10)
}

func main() {
	//channelDemo()

	// 设置缓冲区的channel
	//bufferedChannel()

	// channel close()用法   用range收channel数据
	channelClose()
}

























