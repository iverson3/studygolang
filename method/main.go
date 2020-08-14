package main

import "fmt"

// 方法 和 函数 的区别

// 方法在定义的时候需要与某一变量类型绑定
// 且只有该类型的变量才能调用该方法，方法不能像函数那样直接调用

type Person struct {
	name string
	age int
}
// 定义一个方法，指明调用者的数据类型
func (p Person) say(language string) {
	fmt.Printf("%s can say %s \n", p.name, language)
}

// 调用者可以是个指针类型
func (p *Person) changeName(name string) bool {
	p.name = name
	return true
}

func changeAge(p *Person, age int) bool {
	p.age = age
	return true
}

func main() {
	var p Person
	p.name = "stefan"
	p.age = 22

	// 方法必须要有调用者 (且调用者的变量类型必须跟定义的一样)
	// 且调用者变量会像普通函数参数一样传入方法中
	p.say("english")

	fmt.Println("p.name = ", p.name)

	ok := p.changeName("jack")
	fmt.Println(ok)
	fmt.Println("p.name = ", p.name)


	fmt.Println("p.age = ", p.age)
	changeAge(&p, 27)
	fmt.Println("p.age = ", p.age)

}
