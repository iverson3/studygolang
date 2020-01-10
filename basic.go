package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// go没有全局变量这种说法
// 在函数外定义的变量属于包内"全局"变量
// 注意不要跟函数内部的变量冲突，如果变量命名出现重复 会优先使用函数内部的变量
var aa = 56
var ss = "sssss"
var name = "stefan"  // 与main函数内部的name变量重名

// 使用括号进行简写
var (
	x = 1
	y = "yyy"
	z = false
)

// 定义字符串不能用单引号
//var ss2 = 'fff'

// 在函数外定义变量 不能使用 :=  必须使用var关键字
//ss3 := "sss333"



func variableZeroValue()  {
	var a int
	var s string
	fmt.Println(a, s)
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue()  {
	var a, b int = 12, 5
	var s string = "xxxx"
	fmt.Println(a, b, s)
}

// 类型推断
func variableTypeDeduction()  {
	var a, b, c, s = 3, 6, true, "ccc"
	fmt.Println(a, b, c, s)
}

func variableShorter()  {
	// 定义变量时 简写使用 :=
	a, b, c, s := 3, 6, true, "ccc"
	b = 18
	fmt.Println(a, b, c, s)
}

// 复数 (实部+虚部)
// go内置了对复数的支持
func euler()  {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))

	// 欧拉公式  e的iπ方+1=0
	d := cmplx.Pow(math.E, 1i * math.Pi) + 1
	fmt.Println(d)
	fmt.Printf("%.3f\n", d)
}

func triangle()  {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}

func main() {
	name := "wangfan"
	fmt.Println("abc")

	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()

	fmt.Println(name)

	euler()
	triangle()
}


















