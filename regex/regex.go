package main

import (
	"fmt"
	"regexp"
)

const text = `
my email is ccmouse@gmail.com@www.com
ddddd ffffff@####
eamil2 is abccc@fff.cc
email3 is      kkk@qq.ccc
email4 is  aaa@bbb.com.cn
`

func main() {
	// 创建一个正则表达式
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)

	//match := re.FindString(text)
	//match := re.FindAllString(text, -1)
	match := re.FindAllStringSubmatch(text, -1)  // 子匹配  返回值是二维slice
	fmt.Println(match)

	for _, m := range match {
		fmt.Println(m)
	}
}

















