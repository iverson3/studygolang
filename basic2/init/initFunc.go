package main

import "fmt"
import "studygolang/basic2/init/utils"

// 在golang中，程序的执行顺序是： 全局变量初始化 -> init()函数 -> main()函数

// 定义一个全局变量，并通过调用函数来获得初始值
var age = getAge()

func getAge() int {
	fmt.Println("getAge() function")
	return 20
}

// 初始化函数
// 优先于main()函数执行，进行一些初始化的工作
func init()  {
	fmt.Println("init() function")
}

func main() {
	fmt.Println("main() function， age = ", age)

	fmt.Println("main() function， utils.Name = ", utils.Name, "; utils.Age = ", utils.Age)
}
