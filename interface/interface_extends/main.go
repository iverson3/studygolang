package main

import "fmt"

// 接口与继承的关系比较：
// 当A结构体继承了B结构体，那么A结构体就自动继承了B结构体所有的字段和方法，并且可以直接使用
// 当A结构体需要扩展功能 但不能破坏到原有继承关系时，可以通过实现对应的接口即可；故实现接口可以看作是对继承机制的一种补充

// 接口和继承各自的优势和特点：
// 继承的价值主要在于，解决代码的复用性和可维护性 (代码的高内聚)
// 接口的价值主要在于，设计，设计好各种规范(方法)，让其它自定义类型去实现这些方法
// 1. 接口比继承更加灵活 (继承是 is - a的关系； 而接口是 like - a的关系)
// 2. 接口在一定程度上实现了代码的解耦 (代码的松耦合)


type Person struct {
	Name string
	Age int
}

// 定义一个接口，要求大学生和篮球运动员学习英语技能
type EnglishSkill interface {
	learnEnglish(level int) bool
}

// 学生
type Student struct {
	Person
	Score float64
}
func (s Student) Study(lesson string) {
	fmt.Printf("%s study in %s \n", s.Name, lesson)
}
// 运动员
type Player struct {
	Person
	DressNumber int
}
func (p Player) Exercise(hours float64) {
	fmt.Printf("%s exercise for %.1f hours \n", p.Name, hours)
}

type CollegeStudent struct {
	Student
}
type HighSchoolStudent struct {
	Student
}
func (cs CollegeStudent) learnEnglish(level int) bool {
	fmt.Printf("CollegeStudent %s is learning english in %d level \n", cs.Name, level)
	return true
}

type BasketballPlayer struct {
	Player
}
type FootballPlayer struct {
	Player
}
func (bp BasketballPlayer) learnEnglish(level int) bool {
	fmt.Printf("BasketballPlayer %s is learning english in %d level \n", bp.Name, level)
	return true
}

func main() {
	cs := CollegeStudent{Student{
		Person: Person{
			Name: "jake",
			Age: 23,
		},
		Score:  86,
	}}

	fmt.Println(cs)
	cs.Study("数学")
	cs.learnEnglish(4)


	bp := BasketballPlayer{Player{
		Person: Person{
			Name: "meisi",
			Age: 26,
		},
		DressNumber: 3,
	}}
	fmt.Println(bp)
	bp.Exercise(8)
	bp.learnEnglish(6)
}
