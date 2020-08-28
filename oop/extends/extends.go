package main

import (
	"fmt"
	"math/rand"
)

// golang支持 继承
// 继承可以一定程度的解决代码复用的问题

// 结构体可以使用嵌套匿名结构体中的所有字段和方法(无论字段和方法的首字母是大小写)
// 当结构体和嵌套的匿名结构体有相同的字段或方法时，编译器采用就近访问原则；如希望访问匿名结构体的字段或方法，可通过匿名结构体名来区分
// 当结构体内嵌入多个匿名结构体时，如匿名结构体之间有相同的字段或方法，在访问其相同字段时 必须明确指定匿名结构体的名字进行访问

// 如果结构体内嵌入一个有名结构体，这种模式就是组合；如果是组合关系，在访问组合结构体的字段或方法时，必须使用组合结构体的名字
// 嵌套匿名结构体后，在创建结构体实例时，可以同时初始化各个匿名结构体中的字段

type Student struct {
	Name string
	Age int
	score float64
}
func (s *Student) getScore() float64 {
	return s.score
}
func (s *Student) setInfo(name string, age int)  {
	s.Name = name
	s.Age = age
}

// 定义高中生和大学生 并继承Student结构体
type HighSchool struct {
	Student
	selfField1 string
	selfField2 int
}
func (h *HighSchool) test() {
	fmt.Printf("%s 在进行高中考试 \n", h.Name)
	h.score = float64(rand.Intn(100))
}

type University struct {
	Student
	selfField1 bool
}
func (u *University) test() {
	fmt.Printf("%s 在进行大学考试 \n", u.Name)
	u.score = float64(rand.Intn(150))
}


// 多重继承
// 为了保证代码的简洁性，尽量不要使用多重继承
type Goods struct {
	Name string
	Price float64
}
type Brand struct {
	Name string
	Address string
}
type Tv struct {
	Goods
	Brand
}
type Tv2 struct {
	*Goods
	*Brand
}
// 有名形式的继承，在访问嵌入结构体的字段或方法时，必须使用嵌入结构体的别名进行调用
type Tv3 struct {
	g1 Goods
	g2 Goods
	b1 Brand
}

func main() {
	var h HighSchool
	var u University

	h.setInfo("stefan", 27)
	h.test()
	score := h.getScore()
	fmt.Println("score: ", score)
	fmt.Printf("h: %v \n", h)

	u.setInfo("james", 25)
	u.test()
	score = u.getScore()
	fmt.Println("score: ", score)
	fmt.Printf("u: %v \n", u)


	// 多重继承
	tv1 := Tv{
		Goods: Goods{"名字1", 3258.5},
		Brand: Brand{"品牌1", "天津"},
	}
	fmt.Printf("tv1: %v \n", tv1)
	fmt.Printf("tv1.brand.name: %v \n", tv1.Brand.Name) // 必须明确指定结构体名
	fmt.Printf("tv1.address: %v \n", tv1.Address)

	tv2 := Tv2{
		Goods: &Goods{"名字1", 3258.5},
		Brand: &Brand{"品牌1", "天津"},
	}
	fmt.Printf("tv2: %v \n", tv2)
	fmt.Printf("tv2.goods.name: %v \n", tv2.Goods.Name) // 必须明确指定结构体名
	fmt.Printf("tv2.address: %v \n", tv2.Address)

	tv3 := Tv3{
		g1: Goods{"名字1", 3258.5},
		g2: Goods{"名字2", 6863.8},
		b1: Brand{"品牌1", "天津"},
	}
	fmt.Printf("tv3: %v \n", tv3)
	fmt.Printf("tv3.address: %v \n", tv3.b1.Address)
	fmt.Printf("tv3.name: %v \n", tv3.g2.Name)

}
