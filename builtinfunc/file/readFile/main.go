package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"unicode"
)

// 读取文件内容

type CharCount struct {
	EnCount int
	ZhCount int
	NumCount int
	SpaceCount int
	OtherCount int
}

// 统计字符数
func StatisticsCharCount(filePath string) (CharCount, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return CharCount{}, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var count CharCount
	for {
		str, err := reader.ReadString('\n')

		// 为了兼容中文，需要将字符串转为 []rune 切片类型
		runes := []rune(str)
		for _, c := range runes {
			//fmt.Printf("< %s > \n", string(c))
			switch {
			case c >= 'a' && c <= 'z':
				fallthrough
			case c >= 'A' && c <= 'Z':
				count.EnCount++
			case unicode.Is(unicode.Scripts["Han"], c): // 判断中文字符的方法
			//case unicode.Is(unicode.Han, c):
				count.ZhCount++
			case c >= '0' && c <= '9':
				count.NumCount++
			case c == ' ' || c == '\t':
				count.SpaceCount++
			default:
				count.OtherCount++
			}
		}
		// 读取到文件末尾则退出循环
		if err == io.EOF {
			break
		}
	}

	return count, nil
}

func main() {
	file, err := os.Open("fib.txt")
	if err != nil {
		fmt.Println("open file failed! fail reason: ", err)
	}
	defer file.Close()


	// bufio 带缓冲的读取文件： 分多次，每次只读取一部分文件内容到内存中，直到读取结束 (适合读取大文件)
	reader := bufio.NewReader(file)
	for {
		// ReadString读取直到第一次遇到delim字节，返回一个包含已读取的数据和delim字节的字符串
		str, err := reader.ReadString('\n') // 这里使用"\n"换行符作为参数，即表示一行一行的读取文件内容
		if err == io.EOF { // 读取到文件末尾，则退出for循环
			break
		}
		fmt.Print(str)
	}
	fmt.Println("read file end!")


	// ioutil 一次性的读取文件： 一次性将文件所有内容全部读入到内存中，简单快捷 (适合读取小文件，不需要关心文件的打开和关闭)
	contents, err := ioutil.ReadFile("fib.txt")
	if err != nil {
		fmt.Println("read file failed! reason: ", err)
	}
	fmt.Printf("file contents: \n%s", contents)



	count, err := StatisticsCharCount("fib2.txt")
	if err != nil {
		fmt.Println("字符统计失败，error: ", err)
	}
	fmt.Printf("result: %v", count)
}
