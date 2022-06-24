package main

import (
	"fmt"
	"unsafe"
)

type Duct struct {
	//a int64
	a string
	b int64
}

func main() {

	//test1()


	var str = "abcdefg"

	bytes := *((*[]byte)(unsafe.Pointer(&str)))

	fmt.Println(string(bytes[0]))
	fmt.Println(string(bytes[1]))
	fmt.Println(string(bytes[2]))
}


func test1() {
	var a int64 = 8

	//duct := (*Duct)(unsafe.Pointer(uintptr(unsafe.Pointer(&a)) + 3))
	duct := (*Duct)(unsafe.Pointer(uintptr(unsafe.Pointer(&a))))

	fmt.Println(duct.a)
	fmt.Println(duct.b)
}