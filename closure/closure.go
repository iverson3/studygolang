package main

import (
	"fmt"
	"strings"
)

// 闭包
// 匿名函数自身和其引用到的相关变量 共同组成的整体叫闭包

func makeSuffix(suffix string) func(string) string {
	return func(filename string) string {
		if strings.HasSuffix(filename, suffix) {
			return filename
		} else {
			return filename + suffix
		}
	}
}

func main() {
	f := makeSuffix(".png")
	res := f("abc.jpg")
	fmt.Println("filename = ", res)
}
