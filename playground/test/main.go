package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	// 输入      通过？  正确输出
	// -1        ok     -1
	// 1         ok     -1
	// 10        ok     -1
	// 11        ok     -1
	// 110       ok     -1
	// 101       ok     110
	// 1234      ok     1243
	// 4321      ok     -1
	// 1433      ok     3134
	// 11145     ok     11154
	// 11137765  ok     11153677
	// 1485421   ok     1512448

	num := nextMaxNum(1485421)
	fmt.Println(num)
}

func nextMaxNum(num int64) int64 {
	// 返回-1 表示参数不符合要求或没有满足条件的结果
	if num < 10 {
		return -1
	}
	str := strconv.FormatInt(num, 10)
	pos := 0
	var res int64
	passNums := make([]int, 0, len(str))

	for i := len(str) - 1; i > 0; i-- {
		int1, _ := strconv.ParseInt(string(str[i]), 10, 32)
		passNums = append(passNums, int(int1))
		if str[i] > str[i-1] {
			pos = i
			break
		}
		if i == 1 {
			return -1
		}
	}

	if len(passNums) > 1 {
		sort.Ints(passNums)
	}

	if len(passNums) == 1 {
		newNumStr := str[:pos-1] + string(str[pos]) + string(str[pos-1]) + str[pos+1:]
		res, _ = strconv.ParseInt(newNumStr, 10, 64)
	} else {
		// str中第 pos-1 位的数字
		int2, _ := strconv.ParseInt(string(str[pos-1]), 10, 32)
		replacePos := 0
		// 遍历找到passNums中比pos-1位的数字大 但是在切片中最小的数字
		for i := 0; i < len(passNums); i++ {
			if passNums[i] > int(int2) {
				replacePos = i
				break
			}
		}
		newNumStr := str[:pos-1] + strconv.Itoa(passNums[replacePos])

		hasUsed := false
		for i := 0; i < len(passNums); i++ {
			if i != replacePos {
				if int2 < int64(passNums[i]) && !hasUsed {
					newNumStr = newNumStr + strconv.Itoa(int(int2))
					hasUsed = true
					i--
				} else {
					newNumStr = newNumStr + strconv.Itoa(passNums[i])
				}
			}
		}
		res, _ = strconv.ParseInt(newNumStr, 10, 64)
	}
	return res
}
