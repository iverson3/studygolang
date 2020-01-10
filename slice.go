package main

import "fmt"

// 参数s 类型为[]int 即slice切片
func updateSlice(s []int)  {
	s[0] = 100
}

// 切片
// slice本身并没有数据，它是对底层array的一个view (修改slice就会修改到对应的array)
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:] = ", arr[:])

	s1 := arr[:6]
	s2 := arr[2:]
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	// 在slice上再进行slice
	s3 := arr[:]
	fmt.Println(s3)
	fmt.Println("reSlice")
	s3 = s3[:6]
	fmt.Println(s3)
	s3 = s3[2:]
	fmt.Println(s3)

	// 拓展slice
	// 在底层array的范围内 slice是可以向后扩展的 但不可以向前扩展
	// s[i]不可以超越len(s) 即对slice本身的访问 不允许超过其自身的长度
	// 向后扩展不可以超越底层数组cap(s) 即在reSlice时不能超过底层数组的最大长度
	fmt.Println("Extending slice")
	arr[0], arr[2] = 0, 2
	fmt.Println(arr)

	s4 := arr[2:6]
	s5 := s4[3:5]   // reSlice操作
	fmt.Printf("s4=%v, len(s4)=%d, cap(s4)=%d\n", s4, len(s4), cap(s4))
	fmt.Printf("s5=%v, len(s5)=%d, cap(s5)=%d\n", s5, len(s5), cap(s5))

}

















