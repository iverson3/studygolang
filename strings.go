package main

import (
	"fmt"
	"unicode/utf8"
)

// len()函数获得的是字符串的字节数 (不是字符数)
// 因为一个中文字符通常占3个字节(byte) 所以如果字符串含中文，可能len()的结果不是我们想要的

// []byte(s) 是将字符串转为字节数组(byte数组)  同样的，如果字符串含中文，那么一个中文字符就会被转为3个byte
// []rune(s) 则会自动区分英文字符、中文字符或其他字符，将他们进行对应的处理 最终返回我们需要的字符数组 (而不是字节数组)
// []rune(s) 会自动处理例如中文字符这样的多字节的字符 (英文字母是单字节的字符 一个英文字符只占一个字节)


// 字符串的常见处理函数
// strings.Fields()
// strings.HasPrefix()
// strings.HasSuffix()
// strings.Split()
// strings.Join()
// strings.Repeat()
// strings.Replace()
// strings.Contains()
// strings.Index()
// strings.ToUpper()
// strings.Trim()

func main() {
	s := "Yes我爱慕课网!"

	fmt.Println(len(s))
	fmt.Println([]byte(s))

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s {
		// ch is a rune (int32)
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune Count: ",
		utf8.RuneCountInString(s))

	bs := []byte(s)
	for len(bs) > 0 {
		ch, size := utf8.DecodeRune(bs)
		bs = bs[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()

}
