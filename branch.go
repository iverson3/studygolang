package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	g := ""
	// switch后面可以没有表达式  在case后面加条件判断
	// case里面不用加break  进入某个case后程序会自动break
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		panic(fmt.Sprintf("Wrong score: %d", score))
	}
	return g
}

func main() {
	const filename = "abc.txt"

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		//fmt.Printf("%s\n", bytes)
		fmt.Println(string(bytes))
	}

	g := grade(88)
	fmt.Println(g)

	g = grade(101)
	fmt.Println(g)
}
