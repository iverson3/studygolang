package main

import "fmt"

var arr [3]int
var s1 []int

func main() {
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	s1 = append(s1, 1)
	s1 = append(s1, 2)
	s1 = append(s1, 3)

	s2 := arr[:2]
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
	worker2(s2)
	fmt.Println(s2[1])
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
	fmt.Println()

	worker(arr)
	fmt.Println(arr[1])
	fmt.Println(len(arr))

	worker2(s1)
	fmt.Println(s1[1])
	fmt.Println(len(s1))
}

func worker(data [3]int)  {
	data[1] = 12138
}
func worker2(data []int)  {
	data[1] = 12138

	data = append(data, 555)
	fmt.Printf("worker len: %d \n", len(data))
	fmt.Printf("worker cap: %d \n", cap(data))
}