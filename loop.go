package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// 将十进制的数字转为二进制表示
func convertToBin(n int) string {
	if n == 0 {
		return "0"
	}
	res := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		res = strconv.Itoa(lsb) + res
	}
	return res
}

func printFile(filename string)  {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	printFileContents(file)

	//for i := 1;  i <= 100; i++ {
	//	fmt.Println(i)
	//}

	// 死循环
	//for {
	//	fmt.Println("ever")
	//}
}

func printFileContents(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)
	// 相当于while循环
	for scanner.Scan() {
		// 循环扫描读取文件 并输出每一行的文件内容
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(
		convertToBin(5),  // 101
		convertToBin(13), // 1101
		convertToBin(0),  // 0
		)

	printFile("abc.txt")

	s := `abbbc"d"
 	gggg
	1222222

	pp`

	printFileContents(strings.NewReader(s))
}
