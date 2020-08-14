package main

import (
	"fmt"
	"studygolang/factoryPattern/model"
)

// 使用工厂模式实现跨包创建结构体实例

// 场景：某个包里定义了一个结构体，因为该结构体名的首字母是小写，无法直接在其他包中被引入和使用

func main() {
	// student结构体无法被直接引入
	//var stu = model.student{
	//	Name: "stefan",
	//  age: 22,
	//	Score: 88,
	//}


	// 通过调用工厂函数来获取结构体的实例
	stu := model.NewStudent("stefan", 26, 89)
	stu2 := model.NewStudent2("jake", 42, 77)

	fmt.Printf("stu: %v \n", stu)
	// 因为结构体中的age字段首字母是小写的，所以无法在其他包中进行直接访问
	//fmt.Printf("stu name: %s, age: %d \n", stu.Name, stu.age)
	fmt.Printf("stu name: %s, age: %d \n", stu.Name, stu.GetAge())

	fmt.Printf("stu2: %v \n", stu2)
	fmt.Printf("stu2 name: %s, age: %d \n", stu2.Name, stu2.GetAge())

}
