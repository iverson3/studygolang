package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Person struct {
	name string
	age byte
}

func main() {
	//t1()
	//t2()
	//t3()
	t4()
}

func t1()  {
	var n int64 = 3

	p := unsafe.Pointer(&n)

	*((*int64)(p)) = 5

	fmt.Println(n)
}

func t2()  {
	var a uint8 = 10
	var b uint8 = 20

	change(&a, &b)

	fmt.Println(a)
	fmt.Println(b)
}

func change(a, b *uint8) {
	_ = a
	p := unsafe.Pointer(b)
	// 高位补0
	*((*uint16)(p)) = 3
	// 低位全为0
	*((*uint16)(p)) = 1<<11
}

func t3() {
	me := Person{
		name: "stefan",
		age:  28,
	}

	fmt.Println(unsafe.Alignof(me))
	fmt.Println(unsafe.Sizeof(me.name))
	fmt.Println(unsafe.Sizeof(me.age))
	fmt.Println(unsafe.Sizeof(me))

	p := unsafe.Pointer(&me)
	p2 := unsafe.Pointer(uintptr(p) + unsafe.Offsetof(me.age))

	*((*byte)(p2)) = 1

	fmt.Println(me.name)
	fmt.Println(me.age)
}

func t4() {
	decryptContent := "/AfJDFvdfsadifwhs564hnsa/dFD"

	iv := decryptContent[0:16]
	key := decryptContent[2:18]

	fmt.Println(&iv)
	fmt.Println(&key)
	fmt.Println(unsafe.Sizeof(iv))
	fmt.Println(unsafe.Sizeof(key))

	ivHeader := (*reflect.StringHeader)(unsafe.Pointer(&iv))
	keyHeader := (*reflect.StringHeader)(unsafe.Pointer(&key))
	fmt.Println(ivHeader.Data)
	fmt.Println(keyHeader.Data)

	ivBytes := stringToByte(&iv)
	//keyBytes := stringToByte(&key)

	fmt.Println(string(ivBytes))
	//fmt.Println(string(keyBytes))
}

func stringToByte(key *string) []byte {
	strPtr := (*reflect.SliceHeader)(unsafe.Pointer(key))
	fmt.Println("==================")
	fmt.Println(strPtr.Data)
	fmt.Println(strPtr.Len)
	fmt.Println(strPtr.Cap)

	// 这个赋值将会导致父函数中变量key的Data字段值被覆盖(本来是个地址，被覆盖为strPtr.Len对应的数值)
	//strPtr.Cap = strPtr.Len

	capPtr := unsafe.Pointer(uintptr(unsafe.Pointer(strPtr)) + unsafe.Offsetof(strPtr.Cap))
	str := *(*string)(capPtr)
	fmt.Println("str = ", str)

	return *(*[]byte)(unsafe.Pointer(strPtr))
}