package main

import "fmt"

type Cat struct {
	Name string
	Age int
}

func closeChannel() {
	intChan := make(chan int, 3)

	intChan <- 1
	intChan <- 2

	close(intChan)

	// channel关闭之后，只能取 不能存
	//intChan <- 3

	a := <-intChan
	fmt.Println("a = ", a)
	b := <-intChan
	fmt.Println("b = ", b)
}

// 遍历channel
func forRangeChannel() {
	intChan := make(chan int, 8)
	for i := 0; i < 8; i++ {
		intChan <- i + 1
	}
	close(intChan)

	// 用for range遍历channel时，当遍历完(取完)channel中所有的数据，如果channel已经被关闭，则结束遍历并退出for循环
	// 反之，如果遍历完了 而channel还未关闭，遍历就会继续并等待channel中继续被放入数据；如果之后没有程序再向channel中放入数据，则会导致死锁 程序panic
	for v := range intChan {
		fmt.Println(v)
	}
}

func main() {
	// 定义一个可以传输任意数据类型的channel
	chan1 := make(chan interface{}, 3)

	chan1 <- 10
	chan1 <- "jack"

	cat := Cat{
		Name: "tom",
		Age:  2,
	}
	chan1 <- cat

	// 从channel中取出数据并丢弃
	<-chan1
	<-chan1

	cat1 := <-chan1

	fmt.Printf("cat1 type: %T, value:%v \n", cat1, cat1)
	// 从channel中取出的struct无法直接使用，必须使用类型断言处理
	//fmt.Println("cat.name = ", cat1.Name)

	// 使用类型断言转成对应的数据类型再使用
	cat2 := cat1.(Cat)
	fmt.Println("cat.name = ", cat2.Name)

	closeChannel()

	forRangeChannel()
}
