package main

import (
	"fmt"
	"reflect"
)

func modifyValue(val interface{}) {
	rValue := reflect.ValueOf(val)

	fmt.Println("rValue kind: ", rValue.Kind())

	// 报错，rValue是一个指针类型
	// rValue.SetInt(200)

	// 必须先调用Elem()得到rValue指针指向的值的Value封装 才能去调用set相关方法 修改值
	// Elem()返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装
	rValue.Elem().SetInt(200)
}

func main() {
	num := 10

	modifyValue(&num)

	fmt.Println(num)
	
	
	str := "xxx"
	fs := reflect.ValueOf(&str)

	kind := fs.Kind()
	fmt.Println(kind)

	fs.Elem().SetString("bbb")

	fmt.Println(fs)
	fmt.Println(str)
}
