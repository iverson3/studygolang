package main

import (
	"fmt"
	"strconv"
)

// 基本变量类型的转换
// 必须显示的进行转换，无法自动转换

var f1 float32 = 12.58

func main() {
	i1 := int(f1)

	fmt.Printf("i1 = %v \n", i1)
	fmt.Printf("i1 type is %T \n", i1)
	fmt.Printf("f1 type is %T \n", f1)

	// 转换时需要考虑变量的范围，否则可能出现非预期的结果
	var i2 int64 = 9999999
	var i3 int8 = int8(i2)
	fmt.Println("i3 = ", i3)


	var n1 int32 = 20
	var n2 int64
	n2 = int64(n1) + 10
	fmt.Println("n2 = ", n2)

	var n3 int32 = 15
	var n4 int8
	//var n5 int8
	n4 = int8(n3) + 127
	//n5 = int8(n3) + 128 // 编译不通过，超出了int8的值范围
	fmt.Println("n4 = ", n4) // 不是预期的结果，因为运算结果超出了int8的值范围


	// 基本数据类型转string类型
	var num1 int = 26
	var num2 float32 = 12.678
	var b1 bool = true
	var char1 byte = 'k'
	var str1 string

	// 使用fmt.Sprintf()方法进行转换
	str1 = fmt.Sprintf("%d", num1)
	fmt.Printf("str type: %T, str = %v \n", str1, str1)
	str1 = fmt.Sprintf("%f", num2)
	fmt.Printf("str type: %T, str = %q \n", str1, str1)
	str1 = fmt.Sprintf("%t", b1)
	fmt.Printf("str type: %T, str = %q \n", str1, str1)
	str1 = fmt.Sprintf("%c", char1)
	fmt.Printf("str type: %T, str = %q \n", str1, str1)

	// 使用strconv包里的系列函数来进行转换
	str1 = strconv.FormatInt(int64(num1), 10)
	fmt.Printf("str type: %T, str = %v \n", str1, str1)
	str1 = strconv.FormatFloat(float64(num2), 'f', 10, 64)
	fmt.Printf("str type: %T, str = %q \n", str1, str1)
	str1 = strconv.FormatBool(b1)
	fmt.Printf("str type: %T, str = %q \n", str1, str1)



	// string转为其他基本数据类型
	// 使用strconv包的系列函数
	// 如果string无法转换为对应的类型，则会返回该类型对应的默认值 (比如 "abc"转int 则会返回0  "abc"转float 则会返回false)
	var str2 string = "12138"
	var str3 string = "3.14159"
	var int1 int64
	var f2 float64
	int1, err := strconv.ParseInt(str2, 10, 32)   // 返回值是int64类型
	if err != nil {
		panic(err)
	}
	fmt.Printf("int1 type: %T, int1 = %v \n", int1, int1)

	f2, err = strconv.ParseFloat(str3, 64)  // 返回值是float64类型
	if err != nil {
		panic(err)
	}
	fmt.Printf("f2 type: %T, f2 = %v \n", f2, f2)
}
