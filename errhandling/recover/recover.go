package main

import (
	"fmt"
)

// panic
// 停止当前函数执行，一直向上返回，同时执行每一层的defer；如果没有遇到recover 则程序退出

// recover
// 仅在defer调用中使用，可以获取panic的值，如果无法处理 可重新panic

func tryRecover()  {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(r)
		}
	}()

	//panic(errors.New("this is an error"))

	a := 0
	b := 5 / a
	fmt.Println(b)

	panic(123)
}

func main() {
	tryRecover()
}
