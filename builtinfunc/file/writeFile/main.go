package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 拷贝文件 - 图片、音频、视频都可以进行拷贝
func CopyFile(srcPath string, dstPath string) (written int64, err error) {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		fmt.Println("open src File failed! error: ", err)
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dstPath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open dst File failed! error: ", err)
	}
	defer dstFile.Close()
	writer := bufio.NewWriter(dstFile)

	// 构建 Writer 和 Reader 最后再调用系统的io.Copy()方法即可
	return io.Copy(writer, reader)
}

func main() {
	filePath := "fib2.txt"

	file, err := os.OpenFile(filePath, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file failed! error: ", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		_, _ = writer.WriteString("write string by bufio\r\n")  // "\r\n" 换行字符
	}
	// 因为bufio的writer是带缓冲的，因此在调用WriteString()时，内容是先被写入到缓冲中
	// 需要通过调用Flush()将缓冲中已写入的所有数据一次性写入到文件中
	err = writer.Flush()
	if err != nil {
		fmt.Println("write file failed! error: ", err)
	}

	fmt.Println("write file success!")


	ok, err := PathExist("fsfsd.xxx")
	if ok {
		fmt.Println("file exists.")
	} else {
		if err == nil {
			fmt.Println("file dose not exist.")
		} else {
			fmt.Println("other error: ", err)
		}
	}

}
