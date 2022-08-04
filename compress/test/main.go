package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

const (
	WindowSize = 1024   // 滑动窗口的长度（字节）
    startSearchIndex = 8
    compressMinLength = 8
	maxRepeatLen = 256
)

type Pos struct {
	Index int
	Length int
}

// RepeatBytesInfo 重复字节序的信息
type RepeatBytesInfo struct {
	Index int       // 重复字节序中第一个字节所在的下标
	Distance int    // 重复字节序与跟它重复的字节序的距离
	Length int      // 重复字节序的长度
}

func main() {
	inFile := "./InstallDocker.msi"
	outFile := "./out.txt"
	err := LZ77(inFile, outFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	//outFile2 := "./out2.txt"
	//err = deLZ77(outFile, outFile2)
	//if err != nil {
	//	fmt.Println()
	//}
}

// LZ77 ZIP中的LZ77算法 - 模拟实现
func LZ77(inFileName string, outFileName string) error {
	inFile, err := os.Open(inFileName)
	if err != nil {
		return err
	}
	defer inFile.Close()

	outFile, err := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer outFile.Close()

	inBytes, err := ioutil.ReadAll(inFile)
	if err != nil {
		return err
	}

	// 滑动窗口的左右边界下标
	winL := 0
	winR := startSearchIndex - 1
	// 当前寻找到的位置 （从3开始是因为如果只有1个或2个字节重复，没有压缩的必要）
	searchIndex := startSearchIndex
	// 是否在窗口内找到重复的字节序
	find := false
	// 内容的字节长度
	contentLen := len(inBytes)
	// 记录所有的重复字节序信息
	result := make([]RepeatBytesInfo, 0)

	for {
		find = false
		// 用来存放查找过程中找到的重复字节序
		repeatBytes := make([]byte, 0)
		// 记录最大重复字节序所在的位置信息 (包括位置下标和重复字节序的长度)
		var maxRepeatPos *Pos

		// 在窗口内遍历
		// 在窗口内从右往左搜索重复字节序
		for i := winR; i >= winL ; i-- {
			tmpIndex := searchIndex

			// 如果窗口中剩余的搜索长度小于或等于已经找到的重复字节序的长度，则停止搜索
			if maxRepeatPos != nil && (i - winL) < maxRepeatPos.Length {
				break
			}

			// 寻找第一个相等的字节
			if inBytes[tmpIndex] == inBytes[i] {
				repeatBytes = append(repeatBytes, inBytes[i])

				// 从第一个相等的字节开始，从左往右遍历，比较并将相等的字节append到repeatBytes中；如果遇到有字节不相等则停止
				for j := i + 1; j <= winR; j++ {
					tmpIndex++
					if tmpIndex >= contentLen || len(repeatBytes) >= maxRepeatLen {
						break
					}
					if inBytes[tmpIndex] == inBytes[j] {
						repeatBytes = append(repeatBytes, inBytes[j])
					} else {
						// 一但遇到字节不相等则退出
						break
					}
				}

				// 重复字节序长度大于等于compressMinLength才进行处理
				if len(repeatBytes) >= compressMinLength {
					if maxRepeatPos == nil {
						maxRepeatPos = &Pos{
							Index:  i,
							Length: len(repeatBytes),
						}
					}
					// 只记录最长的重复字节序
					if maxRepeatPos != nil && len(repeatBytes) > maxRepeatPos.Length {
						maxRepeatPos.Index = i
						maxRepeatPos.Length = len(repeatBytes)
					}
				}
				// 重置
				repeatBytes = []byte{}
			}
		}

		if maxRepeatPos != nil {
			result = append(result, RepeatBytesInfo{
				Index: searchIndex,
				Distance: searchIndex - maxRepeatPos.Index,
				Length:   maxRepeatPos.Length,
			})

			find = true
			// 找到则根据找到的字节序长度 往前移动对应的字节数
			searchIndex = searchIndex + maxRepeatPos.Length
		}

		// 没有找到则只需要往进移动一个字节
		if !find {
			searchIndex++
		}
		// 窗口右移
		winR = searchIndex - 1
		if winR - WindowSize >= 0 {
			winL = winR - WindowSize + 1
		}

		// 当窗口移动到最右边，则退出搜索
		if winR == contentLen - 1 {
			break
		}
	}


	fmt.Println(result)

	posMap := make(map[int]RepeatBytesInfo)
	for _, pos := range result {
		posMap[pos.Index] = pos
	}

	outBytes := make([]byte, 0)
	for i := 0; i < contentLen; {
		if pos, ok := posMap[i]; ok {
			distance := []byte(strconv.Itoa(pos.Distance))
			lens := []byte(strconv.Itoa(pos.Length))

			outBytes = append(outBytes, '(')
			outBytes = append(outBytes, distance...)
			outBytes = append(outBytes, ',')
			outBytes = append(outBytes, lens...)
			outBytes = append(outBytes, ')')
			i += pos.Length
		} else {
			outBytes = append(outBytes, inBytes[i])
			i++
		}
	}

	n, err := outFile.Write(outBytes)
	if err != nil || n != len(outBytes) {
		return err
	}

	//fmt.Println(string(inBytes))
	//fmt.Println(string(outBytes))

	return nil
}

// LZ77对应的解码
func deLZ77(inFileName, outFileName string) error {
	inFile, err := os.Open(inFileName)
	if err != nil {
		return err
	}
	defer inFile.Close()

	outFile, err := os.OpenFile(outFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer outFile.Close()

	inBytes, err := ioutil.ReadAll(inFile)
	if err != nil {
		return err
	}
	
	outBytes := make([]byte, 0)
	var tmpI int

	var i = 0
	for {
		if inBytes[i] == '(' {
			tmpI = i
			var j = i + 1
			var distance, length int
			var x = j
			var lens = 1

			for {
				lens++
				_, err = strconv.Atoi(string(inBytes[j]))
				if err != nil {
					if j == i + 1 {
						panic("压缩文件格式不对")
					} else {
						if inBytes[j] == ',' {
							distance, _ = strconv.Atoi(string(inBytes[x:j]))
							j++
							x = j
							continue
						} else if inBytes[j] == ')'{
							length, _ = strconv.Atoi(string(inBytes[x:j]))
							i = j - lens + length + 1
							break
						} else {
							panic("压缩文件格式不对")
						}
					}
				}
				if j >= len(inBytes) {
					panic("压缩文件格式不对")
				}
				j++
			}

			index := tmpI - distance
			if index < 0 || index >= len(inBytes) || index+length >= len(inBytes) {
				panic("压缩文件格式不对")
			}

			outBytes = append(outBytes, inBytes[index:index+length]...)

			leftBytes := make([]byte, len(inBytes) - (tmpI+lens))
			copy(leftBytes, inBytes[tmpI+lens:])

			inBytes = append(inBytes[:tmpI], inBytes[index:index+length]...)
			inBytes = append(inBytes, leftBytes...)
		} else {
			outBytes = append(outBytes, inBytes[i])
			i++
		}
		if i >= len(inBytes) {
			break
		}
	}

	n, err := outFile.Write(outBytes)
	if err != nil || n != len(outBytes) {
		return err
	}

	fmt.Println(string(inBytes))
	fmt.Println(string(outBytes))

	return nil
}
