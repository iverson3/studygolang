package animal

import "fmt"

type Duck struct {
	Color string      // 8+8
	age int64         // 8
	//test int64        // 8
	weight int32      // 4 + 4
	//Name string       // 8+8
}

func NewDuck(color string, age int64, weight int32) Duck {
	return Duck{
		Color:  color,
		age:    age,
		//test: 888,
		weight: weight,
		//Name: "stefan",
	}
}

//func (d Duck) GetAgeOfDuck() int64 {
//	return d.age
//}

func (d Duck) PrintInfo() {
	fmt.Println("duck Info: ", d.Color, d.age, d.weight)
}
