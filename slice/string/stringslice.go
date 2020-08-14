package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "stefan.wang@gmail.com"

	i := strings.IndexByte(str, '@')
	// string底层是一个byte数组，因此可以对string进行切片操作
	s := str[:i]

	fmt.Printf("s type: %T \n", s)  // string类型
	fmt.Println("username = ", s)


	// 字符串是不可以直接进行修改的，需要先将字符串转为 []byte或[]rune类型 再进行修改操作，之后再转为string
	b1 := []byte(str)
	fmt.Printf("b1 type: %T \n", b1)
	fmt.Println("b1 = ", b1)
	fmt.Println("b1[0] = ", b1[0])

	b1[2] = 'a'
	str = string(b1)
	fmt.Println("str = ", str)


	// 注意: 将string转成[]byte后，可以处理英文和数字，但是无法处理中文
	// 因为[]byte类型是按字节来处理的，而一个汉字是3个字节，
	// 解决方法是 将string转成[]rune类型，因为[]rune是按字符来处理的，兼容汉字

	r1 := []rune(str)
	r1[1] = '中'
	str = string(r1)
	fmt.Println("str = ", str)


	// 定义二维切片
	var arr3 [][]int
	fmt.Printf("arr3 type: %T", arr3)
	fmt.Printf("arr3 len: %d", len(arr3))
	fmt.Printf("arr3 cap: %d", cap(arr3))
	arr3 = append(arr3, []int{1})
	arr3 = append(arr3, []int{5, 6, 7})

	fmt.Println()
	fmt.Printf("arr3: %v", arr3)


	var ssss []uint64
	ssss = append(ssss, 6)
	//ssss = append(ssss, -6)
	//ssss = append(ssss, 9999999999999999999999999999999999999999999999999)
	fmt.Println("ssss = ", ssss)



	// 定义二维数组 - 内存分布情况
	var arr4 [2][3]int
	arr4[0][1] = 10

	fmt.Printf("arr4[0] 内存地址: %p \n", &arr4[0])
	fmt.Printf("arr4[1] 内存地址: %p \n", &arr4[1])

	fmt.Printf("arr4[0][0] 内存地址: %p \n", &arr4[0][0])
	fmt.Printf("arr4[1][0] 内存地址: %p \n", &arr4[1][0])
}
