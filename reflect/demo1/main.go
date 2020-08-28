package main

import (
	"fmt"
	"reflect"
)

// reflect 反射

// 反射可以在运行时动态获取变量的各种信息，包括类型(type)和类别(kind)
// 如果是结构体变量，还可以获取到结构体本身的信息(包括结构体的字段和方法)
// 通过反射，可以修改变量的值，可以调用关联的方法

// 下面三种类型可以相互转换，这也是反射能够实现的一个基础
// 某种类型 <---> interface{} <---> reflect.Value

type Person struct {
	Name string
	Age int
}

type Cat struct {
	Name string
	Color string
}

func testReflect(val interface{}) {
	rType := reflect.TypeOf(val)
	fmt.Println("rType = ", rType)

	rValue := reflect.ValueOf(val)
	fmt.Println("rValue = ", rValue)

	// rValue输出好像是个int值，但它的类型并不是int，而是reflect.Value；所以无法进行数学运算
	//num2 := 2 + rValue
	num2 := 2 + rValue.Int() // 可以进行类型转换
	fmt.Println("num2 = ", num2)

	iValue := rValue.Interface()

	num1 := iValue.(int)

	fmt.Printf("num1 type: %T, value: %v \n", num1, num1)
}

func testReflectStruct(val interface{}) {
	rValue := reflect.ValueOf(val)

	kind := rValue.Kind() // 获取kind，比如 Int String Float64 Struct Map Slice Chan Func Array Interface Ptr Uint 等等
	fmt.Println("kind: ", kind)

	// 通过输出rValue的值和类型，好像rValue就是Person结构体的变量(实例)，但去调用结构体的字段和方法是无法通过编译的
	// 因为reflect是运行时的反射，只有在运行时才会明确反射出来的类型，在编译阶段还是不确定的 未知的；所以代码无法通过编译
	// 必须通过类型断言将其转为对应的数据类型，才能进行正常的使用
	// rValue.Name

	// 将reflect.Value类型转为interface{}空接口类型
	iValue := rValue.Interface()

	switch iValue.(type) {
		case Person:
			p := iValue.(Person)
			fmt.Printf("p type: %T; p value: %v \n", p, p)
		case Cat:
			c := iValue.(Cat)
			fmt.Printf("c type: %T; c value: %v \n", c, c)
		default:
			fmt.Println("未知类型")
	}

	// 通过类型断言将空接口类型转为指定的类型
	//p, ok := iValue.(Person)
	//if !ok {
	//	fmt.Println("断言失败")
	//	return
	//}
	//fmt.Printf("p type: %T; p value: %v \n", p, p)

	// 类型转换成功之后可以正常调用结构体的字段和方法
	// p.Name
}

func main() {
	p := Person{
		Name: "jake",
		Age:  34,
	}

	rType := reflect.TypeOf(p)
	fmt.Printf("rType type: %T; rType value: %v \n", rType, rType)

	rValue := reflect.ValueOf(p)
	fmt.Printf("rValue type: %T; rValue value: %v \n", rValue, rValue)

	num1 := 10
	testReflect(num1)

	c := Cat{
		Name:  "tom",
		Color: "yellow",
	}
	testReflectStruct(p)
	testReflectStruct(c)
}
