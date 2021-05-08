package main

import (
	"fmt"
	"runtime"
	"studygolang/slice/trap"
)

// slice切片

func main() {
	data := make([][]int, 0)

	// 如果是服务器，因为需要常驻内存持续提供服务，所以要特别注意内存泄露或无法回收
	for i := 0; i < 100; i++{
		originData := trap.FetchData(128 * 1024)
		res        := trap.DealData(originData)
		//res      := trap.DealData2(originData)

		data = append(data, res)
	}

	//size := unsafe.Sizeof(data)
	//size, err := trap.GetRealSizeOf(data)
	//if err != nil {
	//	fmt.Println("get size failed")
	//} else {
	//	fmt.Println("data mem size: ", size)
	//}

	//var x int
	//size := unsafe.Sizeof(x)
	//fmt.Println(size)

	runtime.GC()

	fmt.Printf("%v\n", data)

	size, _ := trap.GetRealSizeOf(data)
	fmt.Println("data mem size: ", size)

	trap.PrintMem()

	return




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
