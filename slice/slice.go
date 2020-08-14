package main

import "fmt"

// slice切片

func main() {

	arr := []int{1, 2, 3, 4, 5}

	s1 := arr[2:4]

	fmt.Printf("slice1 type: %T \n", s1)
	fmt.Printf("slice1 size: %d \n", len(s1))
	fmt.Printf("slice1 cap: %d \n", cap(s1))

	fmt.Printf("slice1: %v \n", s1)
	fmt.Printf("arr: %v \n", arr)

	//for i, v := range s1 {
	//	fmt.Printf("i = %d, v = %v", i, v)
	//}

	// 通过内置函数为slice追加新的元素
	s1 = append(s1, 100)

	fmt.Printf("slice1: %v \n", s1)
	fmt.Printf("arr: %v \n", arr)

	s1 = append(s1, 500)

	fmt.Printf("slice1: %v \n", s1)
	fmt.Printf("arr: %v \n", arr)

	s2 := []int{7, 8, 9}
	// 可以直接追加一个切片的所有元素
	s1 = append(s1, s2...)

	fmt.Printf("slice1: %v \n", s1)
	fmt.Printf("arr: %v \n", arr)

	// 不能直接追加一个数组的所有元素
	//arr2 := [2]int{6000, 7000}
	//s1 = append(s1, arr2...)


	fmt.Println("============== slice copy ================")

	s3 := []int{1, 2, 3}
	var s4 = make([]int, 8)
	copy(s4, s3)
	fmt.Printf("s4 = %v \n", s4)

	s5 := make([]int, 1)
	copy(s5, s3)
	fmt.Printf("s5 = %v \n", s5)
}
