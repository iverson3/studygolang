package utils

import "fmt"

// 定义两个全局变量
// 供其他包程序使用
var Name string
var Age int

// 对全局变量进行初始化
func init()  {
	Name = "stefan"
	Age = 27
	fmt.Println("utils.go -> init() function")
}
