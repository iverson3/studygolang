package main

import (
	"encoding/json"
	"fmt"
)

// 结构体是值类型
// 结构体中的字段，只有首字母大写，才能在其他包中被访问和使用

// 如果结构体的字段类型是：指针 slice 或 map，它们的默认初始值是 nil
// 在使用这些字段之前 需要使用make()函数分配内存空间 否则无法使用

type Cat struct {
	Name string
	Age int
	Sterilization bool

	hobby []string       // slice
	ptr *int             // 指针
	map1 map[string]int  // map
	func1 func(int) int  // 方法
}

type Point struct {
	x int
	y int
}
type Rect struct {
	leftUp, rightDown *Point
}

// 为字段标tag
type Monster struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	var cat1 Cat

	fmt.Println("cat1: ", cat1)
	fmt.Printf("cat1 addr: %p \n", &cat1)


	cat1.Name = "tom"
	cat1.Age = 2
	cat1.Sterilization = true
	// 定义一个匿名函数赋给func1字段
	cat1.func1 = func(i int) int {
		return i * i
	}

	fmt.Println(cat1.func1(2))

	tmp := 3
	cat1.ptr = &tmp

	// 直接使用会报错，必须先用make()分配空间
	cat1.map1 = make(map[string]int)
	cat1.map1["aaa"] = 111

	cat1.hobby = append(cat1.hobby, "new")

	fmt.Println("cat1: ", cat1)
	fmt.Printf("cat1.Name addr: %p \n", &(cat1.Name))



	// 以下两种创建结构体实例的方法，返回的都是指针
	var cat2 *Cat = new(Cat)
	var cat3 *Cat = &Cat{}

	// 也可以在创建结构体实例的时候直接给字段赋初始值
	var cat4 *Cat = &Cat{
		Name: "marry",
	}

	// 对于结构体指针变量，访问字段的标准方式应该是 (*cat2).Name = "hurry"
	// 但为了使用更加方便，更加符合编程习惯，golang做了简化处理： cat2.Name = "hurry"
	(*cat2).Name = "hurry"
	cat2.Name = "hurry"

	(*cat3).Age = 100
	cat3.Age = 100

	cat4.Age = 18


	// 结构体变量在内存中的分布情况
	// 规律分布，具有连续性
	r1 := Rect{&Point{10, 20}, &Point{3, 4}}

	fmt.Printf("r1 value: %v \n", r1)
	fmt.Printf("r1 addr: %p \n", &r1)
	fmt.Printf("r1.leftUp 本身地址: %p； r1.rightDown 本身地址：%p \n", &r1.leftUp, &r1.rightDown)
	fmt.Printf("r1.leftUp 指向地址: %p； r1.rightDown 指向地址：%p \n", r1.leftUp, r1.rightDown)


	// struct的每个字段上，可以写上一个tag，该tag可以通过反射机制获取，常用的使用场景是 序列化与反序列化
	monster := Monster{
		Name:  "牛魔王",
		Age:   500,
		Skill: "芭蕉扇",
	}
	// 将结构体序列化为json字符串
	marshal, err := json.Marshal(monster)
	if err != nil {
		panic(err)
	}
	fmt.Printf("json.Marshal result: %v", string(marshal))

}
