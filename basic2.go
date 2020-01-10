package main

import (
	"fmt"
	"math"
)

// 常量
// 函数内或包内都可定义常量  定义常量也可使用括号来简化
// go的常量命名不用像其他语言一样 使用全大写字母，直接像普通变量一样正常命名即可
func consts()  {
	const filename = "abc.md"
	// 定义常量的时候 如果没有明确指定常量类型 那么类型是不确定的(相对不确定)
	// 比如下面的常量a b  使用的时候可以当做 int或float
	const a, b = 3, 4
	var c int

	c = b
	fmt.Println(c, b)
	c = int(math.Sqrt(a * a + b * b))

	fmt.Println(filename, c, b)
}

// 枚举类型
func enums()  {
	// 定义枚举类型
	// go没有专门定义枚举的关键字
	const (
		cpp    = 0
		java   = 1
		python = 2
		golang = 3
	)

	// 使用iota定义自增值的枚举类型
	const (
		php = iota
		_
		ruby
		javascript
	)

	// 定义  b kb mb gb tb pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, java, python, golang)
	fmt.Println(php, ruby, javascript)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("main")

	consts()
	enums()
}
