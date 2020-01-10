package main

import "fmt"

func lengthOfNonRepeatingSubStr_old(s string) int {
	// 记录s中每个出现过的字母最后出现的位置索引
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	// []byte(s) 将字符串s转成byte数组(字节数组) 即一个个的字节 (所以目前这种方式只支持英文字符 不支持中文)
	// []rune(s) 将字符串s转成rune数组(字符数组) 即一个个的字符 (支持中文等各种多字节的字符)
	for i, ch := range []rune(s) {
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

// 为了性能考虑 把lastOccurred这个slice放到函数外面定义
// 避免函数多次运行时 同一个slice被多次创建和销毁(费时)
// stores last occurred pos + 1.
// 0 means not seen.
var lastOccurred = make([]int, 0xffff)

// 优化版
func lengthOfNonRepeatingSubStr(s string) int {
	// 还原数据
	for i := range lastOccurred {
		lastOccurred[i] = 0
	}
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		lastI := lastOccurred[ch]
		if lastI > start {
			start = lastI
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i + 1
	}
	return maxLength
}

func main() {
	// 从下面的代码可以看出，将中文字符转为[]byte时 一个中文字会被转成三个byte (而英文字母都是只会转成一个对应的byte)
	fmt.Println([]byte("abc"))
	fmt.Println([]byte("默课网"))

	bs := []byte("默课网")
	bs1 := bs[:3]   // 三个byte为一个中文字
	bs2 := bs[3:6]
	bs3 := bs[6:]
	fmt.Println(string(bs1), string(bs2), string(bs3))
}
