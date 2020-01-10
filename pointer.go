package main

import "fmt"

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
}
