package main

import "fmt"

func main() {
	var hasScanName []int

	t1(hasScanName)

	fmt.Println("=====")
	printIntArr(hasScanName)
}

func t1(hasScanName []int)  {
	t2(hasScanName)
}

func t2(hasScanName []int)  {
	hasScanName = append(hasScanName, 3)

	printIntArr(hasScanName)

	t3(hasScanName)

	printIntArr(hasScanName)
}
func t3(hasScanName []int)  {
	fmt.Println("%%%%%%%%%%%%")
	printIntArr(hasScanName)
	hasScanName = append(hasScanName, 5)
	printIntArr(hasScanName)
	fmt.Println("%%%%%%%%%%%%")
}

func printIntArr(arr []int) {
	for _, v := range arr {
		fmt.Printf(" %d ", v)
	}
	fmt.Println()
}