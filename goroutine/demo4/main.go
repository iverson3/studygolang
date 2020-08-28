package main

import (
	"fmt"
	"time"
)

// channel可以定义属性为只读或只写

// 通过定义形参的类型为 "chan<- int"，该参数在该函数内部为只写的
func Write(in chan<- int) {
	for i := 0; i < 20; i++ {
		in<- i + 1
	}

	// 编译无法通过，不能对只写属性的channel进行读取操作
	// v := <-in
}

// 通过定义形参的类型为 "<-chan int"，该参数在该函数内部为只读的
func Read(out <-chan int) {
	for {
		v, ok := <-out
		if !ok {
			break
		}
		fmt.Println("get number: ", v)
	}

	// 编译无法通过，不能对只读属性的channel进行写入数据的操作
	// out<- 567
}

func main() {
	intChan := make(chan int, 10)

	go Write(intChan)
	go Read(intChan)

	time.Sleep(time.Second * 1)

	fmt.Println("end.")
}
