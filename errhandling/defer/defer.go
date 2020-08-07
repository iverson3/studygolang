package main

import (
	"bufio"
	"fmt"
	"os"
	"studygolang/functional/fib"
)

// defer 使用场景：
// 文件的打开关闭  连接的打开断开  Lock/Unlock
// defer语句遵循堆栈的先入后出规则
// 当defer将语句放入栈中时，也会将相关的变量的值拷贝到栈中，不受后续代码对相关变量修改的影响

func tryDefer()  {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)

	panic("error occurred")
	return

	fmt.Println(4)
}

func writeFile(filename string)  {
	//file, err := os.Create(filename)
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			// 未知类型的错误
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		//panic(err)
		//fmt.Println("Error:", err.Error())
		return
	}
	defer file.Close()

	// 使用bufio向文件中写东西会比较快 (先写入内存 最后再一次性写入文件)
	// writer.Flush() 即把暂存在内存中的数据写入文件中
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		// 将f()函数的返回结果写入到writer buf内存中
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//tryDefer()

	writeFile("fib.txt")
}
