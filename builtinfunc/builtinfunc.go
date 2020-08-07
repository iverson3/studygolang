package main

import "fmt"

// golang系统内建函数
// 程序中可以直接使用的函数，无需引入对应的包

// 比如 len() cap() make() new() append() copy() delete() close() panic() recover()

func main() {
	// 内建函数make()分配并初始化一个类型为切片、映射、或通道的对象。

	// 定义一个int类型的数组 长度为10 (默认值为0)
	s := make([]int, 10)
	fmt.Printf("type: %T, value: %v \n", s, s)

	// 定义一个string类型的channel管道
	chan1 := make(chan string, 0)
	fmt.Printf("type: %T, value: %v, addr: %v \n", chan1, chan1, &chan1)


	// 内建函数new分配内存。其第一个实参为类型，而非值。其返回值为指向该类型的新分配的零值的指针。
	n := new(int)
	*n = 10
	fmt.Printf("type: %T, value: %v, addr: %v \n", n, n, &n)
	fmt.Printf("true value: %v \n", *n)
}
