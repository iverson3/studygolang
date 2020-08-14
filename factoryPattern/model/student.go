package model

// 定义结构体
// 因为结构体名的首字母是小写，所以无法直接在其他包中被引入和使用

type student struct {
	Name string
	age int
	Score float64
}

// 因为age字段是无法在其他包中被直接访问的，所以这里定义一个方法来获取age字段
func (stu student) GetAge() int {
	return stu.age
}

func NewStudent(name string, age int, score float64) student {
	return student{
		Name:  name,
		age: age,
		Score: score,
	}
}

// 返回指针类型
func NewStudent2(name string, age int, score float64) *student {
	return &student{
		Name:  name,
		age: age,
		Score: score,
	}
}