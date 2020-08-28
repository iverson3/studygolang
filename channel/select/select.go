package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 使用select来进行channel的调度
// 在select中使用 nil channel 可以正确运行但不会被select到

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			//time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func createWorker(id int) chan int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func worker(id int, c chan int) {
	// for n := range c {
	// }
	for {
		n := <-c
		// 通过Sleep()降低worker收数据的速度
		//time.Sleep(1 * time.Second)
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)

	// 倒计时： 返回值是一个channel, 在指定的时间之后会向该channel发送一个数据 (只发送一次)
	tm := time.After(10 * time.Second)
	// 定时器： 返回值是一个channel, 每隔指定的时间都会向该channel发送一个数据 (一直间隔的发)
	tick := time.Tick(time.Second)

	// 实际场景中, 发数据和收数据的速度可能是不一样的
	// 创建一个slice 当做queue队列, 存放所有从c1 c2中收到的数据, 再等worker慢慢的从queue中收数据
	var values []int
	for {
		var activeWorker chan int  // nil channel
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]  // 每次从队列中取一个数据，准备发给worker的channel
		}
		// 800毫秒的倒计时, 且每次循环倒计时都会被重置
		timeOutChan := time.After(800 * time.Millisecond)
		// 从两个channel(c1 c2)中收数据 谁先来就从哪里收
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)

		case activeWorker <- activeValue:  // 把从c1 c2中收到的数据 发到worker channel中去
			values = values[1:]

		case <-timeOutChan:  // 收数据的时间间隔超过了800毫秒, 则会走这个case
			fmt.Println("Timeout.")
		case <-tick:
			fmt.Printf("queue len =%d\n", len(values))  // 每隔一段时间打印一下队列的长度
		case <-tm:
			fmt.Print("bye.")
			return
		//default:
		//	fmt.Println("default case------")
		}
	}
}
