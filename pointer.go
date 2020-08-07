package main

import "fmt"

// 引用类型，变量存储的是一个内存地址，这个地址所对应的内存空间才真正存储数据
// 内存通常在堆上分配，当没有任何变量引用这个地址时，该地址对应的数据空间就成为一个垃圾，由GC来回收

// 定义参数a b为指针类型
func swap(a, b *int)  {
	*a, *b = *b, *a
}

func swap2(a, b int) (int, int) {
	return b, a
}

// go 默认都是值传递 (不是引用传递)
// 使用指针可实现引用传递的效果
func main() {
	a, b := 2, 5
	// & 取地址
	swap(&a, &b)
	fmt.Println(a, b)

	//a, b = swap2(a, b)
	//fmt.Println(a, b)


	// 值类型都有其对应的指针类型，写法为 "*数据类型"
	// 值类型包括 int系列 float系列 bool string 数组 结构体struct
	// 引用类型： 指针 slice切片 map 管道chan  interface等
	var n1 int = 10
	fmt.Printf("n1 address: %v \n", &n1)

	// ptr是一个指向int数据类型的指针，ptr的值是一段内存地址
	var ptr *int = &n1
	fmt.Printf("ptr = %v \n", ptr)
	fmt.Printf("指针变量ptr指向的内存地址所存的值是：%v \n", *ptr)

	*ptr = 125
	fmt.Printf("n1 = %d \n", n1)

	t()
}

func t()  {
	var a int = 1
	var b int = 2

	b = a + b
	a = b - a
	b = b - a

	fmt.Println(a)
	fmt.Println(b)
}
