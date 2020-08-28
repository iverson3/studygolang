package main

import (
	"flag"
	"fmt"
)

// 获取 命令行参数

// 方式一： (比较原始 不够灵活)
// os.Args 是一个string切片，存放着所有的命令行参数 (第一个参数是执行文件本身的path)

// 方式二： (推荐且常用)
// flag包解析命令行参数

func main() {
	// 方式一
	//args := os.Args
	//fmt.Println("args: ", args)
	//
	//for i, v := range args {
	//	fmt.Printf("Args[%d]: %s\n", i, v)
	//}


	// 方式二
	var user string
	var pwd string
	var port int
	flag.StringVar(&user, "user", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "888888", "密码，默认为6个8")
	flag.IntVar(&port, "port", 3306, "端口，默认为3306")
	// 必须调用该方法，才能实现参数解析
	flag.Parse()

	fmt.Println("user: ", user)
	fmt.Println("password: ", pwd)
	fmt.Println("port: ", port)
}
