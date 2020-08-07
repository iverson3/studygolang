package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// 函数两个返回值通常是 第一个返回值是函数的结果 第二个返回值是函数运行出错的错误信息
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

// go函数可以有多个返回值
// 实现 13/3 = 4 ... 1
func div(a, b int) (q, r int)  {
	//return a / b, a % b

	q = a / b
	r = a % b
	return
}

// 自定义数据类型 (这里opFuncType是一个自定义的函数类型)
type opFuncType func(int, int) int

// 函数作为函数的参数
func apply(op opFuncType, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, %d)\n", opName, a, b)

	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数列表  (go没有默认参数 可选参数这些)
func sum(numbers ...int) int {
	s := 0
	// 参数numbers相当于一个包含了所有参数的数组
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	res, err := eval(5, 6, "+")
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(res)
	}

	fmt.Println(div( 13, 3))

	// 接收两个返回值 变量命名可随意
	q, r := div(12, 5)
	fmt.Println(q, r)

	// 如果不想要某个返回值 可以用_代替
	a, _ := div(12, 5)
	fmt.Println(a)

	fmt.Println(
		apply(pow, 3, 4),
		)

	// 直接定义匿名函数进行处理
	res2 := apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4)
	fmt.Println(res2)

	fmt.Println(sum(1, 2, 3, 4, 5))
}


















