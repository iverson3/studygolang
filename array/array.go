package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 数组是值类型
// 在函数间传递的时候默认是以值传递的方式进行值拷贝的，如果数组很大，值拷贝就可能比较耗时 费内存
// 此时可以考虑使用引用传递；但注意在使用引用传递时，函数里对数组所做的修改会影响到该数组本身

func test1(arr [3]int)  {
	arr[0] = 1111
}

func test2(arr *[3]int)  {
	fmt.Printf("arr: %v; a: %T; b: %v; c: %v \n", arr, arr, *arr, &arr)
	arr[0] = 1111
}

func test3()  {
	var c1 [26]byte
	for i := 0; i < 26; i++ {
		c1[i] = 'A' + byte(i)
	}

	fmt.Printf("byte: %v \n", c1)

	for i := 0; i < 26; i++ {
		fmt.Printf("%c ", c1[i])
	}
	fmt.Println()
}

func test4()  {
	var arr [5]int

	// 设置随机数的种子，种子变化 生成的随机数才会变化
	// 如果不设置种子或者种子的值不变，那么每次运行生成的随机数都是相同的 不会有变化
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)  // 生成随机数，取值范围在[0,n)的伪随机int值
	}

	fmt.Printf("test4 arr 反转前: %v \n", arr)

	len := len(arr)
	for i := 0; i < len / 2; i++ {
		arr[i], arr[len - 1 - i] = arr[len - 1 - i], arr[i]
	}
	fmt.Printf("test4 arr 反转后: %v \n", arr)
}

func main() {
	// 给定长度，即定义数组
	//var arr1 [10]string
	//var arr2 = make([]string, 10)
	// 不给定长度，即定义slice切片
	//var s1 []string
	//var s2 = make([]string, 0)

	//fmt.Printf("arr: %T \n", s1)
	//fmt.Printf("arr: %T \n", s2)


	arr := [3]int{1, 2, 3}

	//test1(arr)
	test2(&arr)
	test3()
	test4()

	fmt.Printf("arr: %v \n", arr)
}
