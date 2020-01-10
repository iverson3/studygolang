package main

import (
	"awesomeProject1/functional/fib"
	"bufio"
	"fmt"
	"io"
	"strings"
)

// 函数是一等公民, 变量 参数 返回值都可以是函数;
// 函数甚至还可以实现接口 (go的接口实现中 接口的调用者其实就是一个参数)
// 高阶函数
// 函数闭包

// 让intGen函数类型实现Reader接口的Read方法
func (g fib.IntGen) Read(p []byte) (n int, err error) {
	next := g()
	// 当斐波拉契数超过1000后则停止继续生成
	if next > 1000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fib.Fibonacci()

	//fmt.Println(f())   // 1
	//fmt.Println(f())   // 1
	//fmt.Println(f())   // 2
	//fmt.Println(f())   // 3
	//fmt.Println(f())   // 5
	//fmt.Println(f())   // 8
	//fmt.Println(f())   // 13
	//fmt.Println(f())   // 21

	printFileContents(f)

}
