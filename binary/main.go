package main

import (
	"fmt"
	"strconv"
)

// 二进制

func main() {
	test()
	return


	var b1 int = 10

	fmt.Printf("%b\n", b1)

	b2 := strconv.FormatInt(int64(b1), 2)
	fmt.Println(b2)



	b3 := 1<<10 + 1<<8 + 1<<2
	//fmt.Println(b3)
	fmt.Printf("%b\n", b3)

	// 取出高八位
	b6 := b3 >> 8
	fmt.Printf("%b\n", b6)

	// 取出第八位
	fmt.Printf("%b\n", 128)
	b4 := b3 & 128
	fmt.Printf("%b\n", b4)

	// 取出低七位
	fmt.Printf("%b\n", 127)
	b5 := b3 & 127
	fmt.Printf("%b\n", b5)
}

func test()  {
	var c coder = &Gopher{10}

	fmt.Printf("%v \n", c)
	c.code()
	c.debug() // 10
}

type coder interface {
	code()
	debug()
}

type Gopher struct {
	age int
}

func (p Gopher) code() {
	fmt.Printf("%v \n", p)
	p.age += 1
}

func (p *Gopher) debug() {
	fmt.Printf("%v \n", p)
	fmt.Println(p.age)
}
