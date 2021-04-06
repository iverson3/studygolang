package main

import "fmt"

var names []string
var scores []int

var total = 664
var totalLevel = 16

var resultCount = 0


func main() {
	//scores = []int{6, 8, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10,    12, 12, 14, 23, 23, 25,
	//	30, 32, 35, 40, 43, 43, 43, 44, 45, 45, 46, 48, 51, 52, 53, 53, 53, 53, 53, 53, 53, 53, 54, 54, 54, 54}
	//
	//names = []string{"aa", "bb", "cc", "dd", "ee", "ff", "gg", "hh", "ii", "jj", "kk", "ll", "mm", "nn", "oo", "pp",
	//	"qq", "rr", "圣", "涛", "明", "戴", "舒", "丹", "潘", "林", "邵", "伟", "潘", "葛", "秦", "张",
	//	"舒", "涛", "郑", "郑", "周", "董", "丹", "候", "周", "爽", "秦", "胡", "付", "胡", "伟", "培"}


	scores = []int{22, 22, 26, 26, 30, 30, 31, 31, 31, 32, 33, 34, 36, 37, 37, 38, 38, 39, 40, 40, 40, 40,
		43, 43, 43, 43, 44, 44, 44, 45, 45, 46, 46, 48, 51, 52, 53, 53, 53, 53, 53, 53, 53, 53, 54, 54, 54, 54}

	names = []string{"刘", "培", "刘", "芬", "范", "邵", "谢", "谢", "付", "葛", "明", "陈", "焱", "陈", "董", "焱",
		"戴", "范", "李", "涛", "明", "戴", "舒", "丹", "潘", "林", "邵", "伟", "潘", "葛", "秦", "张",
		"舒", "涛", "郑", "郑", "周", "董", "丹", "候", "周", "爽", "秦", "胡", "付", "胡", "伟", "培"}

	//fmt.Println(len(names))
	//fmt.Println(len(scores))
	//return

	var hasScanIndex []int
	var hasScanName []string
	var hasScanScore []int
	worker(0, 0, hasScanIndex, hasScanName, hasScanScore)

	fmt.Println(resultCount)
}

func worker(level int, sum int, hasScanIndex []int, hasScanName []string, hasScanScore []int) {

	//fmt.Println(level)

	if sum > total {
		return
	}
	if level == totalLevel {
		if sum == total {
			resultCount++
			fmt.Printf("===== %d \n", sum)

			printIntArr(hasScanIndex)
			printStrArr(hasScanName)
			printIntArr(hasScanScore)
		}
		return
	} else {
		for i := level; i < len(names); i++ {
			if is_in(hasScanIndex, i) {
				continue
			}
			if is_in_str(hasScanName, names[i]) {
				continue
			}

			if sum + scores[i] > total {
				break
			}

			goWorker(i, level, sum, hasScanIndex, hasScanName, hasScanScore)
		}
	}
}

func goWorker(i int, level int, sum int, hasScanIndex []int, hasScanName []string, hasScanScore []int) {
	sum = sum + scores[i]
	hasScanIndex = append(hasScanIndex, i)
	hasScanName = append(hasScanName, names[i])
	hasScanScore = append(hasScanScore, scores[i])

	//if level > i {
	//	fmt.Printf("%d ------ %d \n", level, i)
	//}

	worker(i, sum, hasScanIndex, hasScanName, hasScanScore)
}

func is_in(arr []int, num int) bool {
	for _, v := range arr {
		if v == num {
			return true
		}
	}
	return false
}
func is_in_str(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func printStrArr(arr []string) {
	for _, v := range arr {
		fmt.Printf(" %s ", v)
	}
	fmt.Println()
}
func printIntArr(arr []int) {
	for _, v := range arr {
		fmt.Printf(" %d ", v)
	}
	fmt.Println()
}