package main

import "fmt"

// 类型断言

type Point struct {
	x, y int
}

func TypeJudge(items ...interface{})  {
	for i, item := range items{
		switch item.(type) {
		//case interface{}:
		//	fmt.Printf("第%d个参数是空接口类型，值是%v \n", i, item)
		case bool:
			fmt.Printf("第%d个参数是bool类型，值是%v \n", i, item)
		case int, int8, int16, int32, int64:
			fmt.Printf("第%d个参数是整型，值是%v \n", i, item)
		case float32, float64:
			fmt.Printf("第%d个参数是浮点类型，值是%v \n", i, item)
		case string:
			fmt.Printf("第%d个参数是string类型，值是%v \n", i, item)
		// 判断自定义的结构体类型
		case Point:
			fmt.Printf("第%d个参数是Point结构体值类型，值是%v \n", i, item)
		case *Point:
			fmt.Printf("第%d个参数是Point结构体引用类型，值是%v \n", i, item)
		default:
			fmt.Printf("第%d个参数的类型不确定，值是%v \n", i, item)
		}
	}
}

func main() {
	var x interface{} // 空接口类型，可以接收任意其他类型
	p := Point{3,4}
	var p2 Point

	// 一个确定类型的变量赋给空接口类型的变量 是可以的
	x = p
	// 一个空接口类型的变量 无法直接赋给确定类型的变量，需使用类型断言
	p2 = x.(Point)

	fmt.Println(p2)


	var f float32 = 2.6
	var f2 float32
	var y interface{}

	y = f
	// 使用类型断言
	f2 = y.(float32)
	// 不使用类型断言，直接使用自动类型推导":="
	f3 := y

	// 下面类型断言出错，因为空接口类型变量y指向的类型是float32，不能使用断言转成float64
	var f4 float64
	//f4 = y.(float64)

	// 类型断言 返回是否成功的结果，不会panic()
	f4, ok := y.(float64)
	if ok {
		fmt.Printf("类型断言成功，f4 type: %T, value: %v \n", f4, f4)
	} else {
		fmt.Println("类型断言失败")
	}

	fmt.Printf("f2 type: %T, value: %v \n", f2, f2)
	fmt.Printf("f3 type: %T, value: %v \n", f3, f3)


	TypeJudge(12.1, 12, "xxx", false, x, Point{}, &Point{})
}
