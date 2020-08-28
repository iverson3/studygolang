package main

import (
	"fmt"
	"time"
)

// 使用select来进行channel的调度
// select可以解决从channel中取数据时的阻塞问题

// 可以通过在goroutine中使用defer+recover捕获程序错误，防止协程中出现的panic导致整个程序奔溃结束运行

func worker(no int, in chan string)  {
	// 通过defer+recover对协程中发生的错误进行捕获并处理，以至于某个协程的奔溃不会导致整个程序的奔溃、不会影响其他协程包括主线程的正常运行
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("worker[%d]协程运行期间出现错误，已退出； error: %v \n", no, err)
		}
	}()

	if no == 3 {
		var map1 map[string]int
		map1["aaa"] = 123 // 这行代码在运行时 会panic，如果不做处理，会导致整个程序奔溃结束
	}
	for i := 0; i < 10; i++ {
		in<- fmt.Sprintf("data = %d", i + 1)

		fmt.Printf("worker[%d] send data: %d \n", no, i + 1)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	dataChan := make(chan string, 40)
	for i := 1; i < 5; i++ {
		go worker(i, dataChan)
	}

	//for {
	//	v, ok := <-dataChan
	//	if !ok {
	//		break
	//	}
	//
	//	fmt.Println("get data: ", v)
	//}

	//for v := range dataChan {
	//	fmt.Println("get data: ", v)
	//}

	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
	}

	fmt.Println("main() end.")
}
