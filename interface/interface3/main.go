package main

import (
	"fmt"
	"reflect"
)

type Person interface {
	Show()
}

type Student struct {}

func (s *Student) Show() {

}

func live() Person {
	var stu *Student
	return stu
}




func main() {

	//a :="aaa"
	//ssh := *(*reflect.StringHeader)(unsafe.Pointer(&a))
	//b := *(*[]byte)(unsafe.Pointer(&ssh))
	//fmt.Printf("%v",b)

	m := 4
	arr := [9]int{1, 3, -1, -3, 5, 3, 6, 7, 5}

	var set []int
	var res []int

	for i := range arr {
		if len(set) > 0 && arr[i] > set[0] {
			set = []int{}
		}
		set = append(set, arr[i])

		if i >= m - 1 {
			res = append(res, set[0])
		}
	}

	fmt.Printf("%v\n", res)

	return


	p := live()

	fmt.Printf("%v\n", p)
	fmt.Println(reflect.TypeOf(p))
	fmt.Println(reflect.ValueOf(p))

	// 接口在判断nil的时候，需要动态类型和动态值都为nil，结果才会为nil
	if p == nil {
		fmt.Println("aaa")
	} else {
		fmt.Println("bbb")
	}



	var s *Student
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.ValueOf(s))

	// s不是接口，判断nil的时候，只需要判断值是否为nil即可
	if s == nil {
		fmt.Println("cccccccc")
	}
}
