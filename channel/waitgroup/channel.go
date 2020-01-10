package main

import (
	"fmt"
	"sync"
)

// goroutine之间的双向通道就是channel
// go slogan: 不要通过共享内存来通信, 要通过通信来共享内存

// channel是一等公民 可以作为参数或返回值

// chan<- int   该channel只能发数据 (向channel里面发数据)
// <-chan int   该channel只能收数据 (从channel里面收数据)

func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n", id, n)
		w.done()
	}
}

type worker struct {
	in chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}

	go doWork(id, w)
	return w
}

// 使用系统提供的 sync.WaitGroup 来实现等待goroutine任务的结束
func channelDemo()  {
	var wg sync.WaitGroup

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i + 1, &wg)
	}

	wg.Add(20)  // 总共有20个异步任务 (也可以在循环里面每次加一个任务  wg.Add(1) )
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	// 程序会一直等在这里，直到所有的goroutine都执行结束并调用 wg.Done()
	wg.Wait()
	fmt.Println("channelDemo end")
}

func main() {
	channelDemo()
	fmt.Println("main end")
}

























