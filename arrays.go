package main

import "fmt"

func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
// 传递数组的引用
func printArray2(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}


// 很多地方可通过下划线 _ 省略变量
// 数组是值类型
// go语言中一般不直接使用数组 而是使用切片
func main() {
	// 定义数组
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}

	var grid [3][4]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	printArray(arr1)
	fmt.Println(arr1)

	printArray2(&arr1)
	fmt.Println(arr1)

	// [3]int 和 [5]int 是不同的类型
	//printArray(arr2)

}
















