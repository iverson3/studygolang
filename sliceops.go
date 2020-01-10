package main

import "fmt"

func printSlice(s []int)  {
	fmt.Printf("s=%v, len(s)=%d, cap(s)=%d\n", s, len(s), cap(s))
}

// slice的各种操作
func main() {
	// 定义数组
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	s4 := arr[2:6]  // 生成slice
	s5 := s4[3:5]   // reSlice操作
	fmt.Printf("s4=%v, len(s4)=%d, cap(s4)=%d\n", s4, len(s4), cap(s4))
	fmt.Printf("s5=%v, len(s5)=%d, cap(s5)=%d\n", s5, len(s5), cap(s5))

	s6 := append(s4, 10)
	s7 := append(s6, 11)
	s8 := append(s7, 12)

	// s6 s7 依旧是arr的一个view
	// s8 却是新数组的view (长度超过arr之后 系统自动开辟的新的数组空间)

	// 对slice进行append操作时，如果长度超过了cap的大小) 系统会自动重新分配一个更大的底层数组来供slice进行view
	// 且新的底层数组的长度是旧slice cap大小的两倍


	fmt.Printf("s6=%v, len(s6)=%d, cap(s6)=%d\n", s6, len(s6), cap(s6))
	fmt.Printf("s7=%v, len(s7)=%d, cap(s7)=%d\n", s7, len(s7), cap(s7))
	fmt.Printf("s8=%v, len(s8)=%d, cap(s8)=%d\n", s8, len(s8), cap(s8))

	fmt.Println("arr =", arr)


	// 直接定义slice (无初始值)
	var s []int

	// slice的cap长度会不断的翻倍 1 2 4 8 16 32 64 128
	for i := 0; i < 100; i++ {
		s = append(s, i)
	}
	printSlice(s)

	// 直接定义slice (有初始值 系统自行计算len和cap)
	s11 := []int{2, 4, 6, 8}
	printSlice(s11)

	// 直接定义slice (不赋初始值[系统自动赋初始值] 但给定len长度)
	s22 := make([]int, 12)
	printSlice(s22)

	// 直接定义slice (不赋初始值[系统自动赋初始值] 但给定len和cap长度)
	s33 := make([]int, 10, 30)
	printSlice(s33)


	// copy slice
	copy(s22, s11)
	printSlice(s22)



	// delete elements from slice

	// 删除s22中第四个元素
	s22 = append(s22[:3], s22[4:]...)
	printSlice(s22)

	// 删除s22第一个元素
	front := s22[0]
	s22 = s22[1:]
	fmt.Println("front =", front)
	printSlice(s22)
	// 删除s22最后一个元素
	tail := s22[len(s22) - 1]
	s22 = s22[:len(s22) - 1]
	fmt.Println("tail =", tail)
	printSlice(s22)
}


























