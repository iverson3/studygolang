package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name string `json:"name"`
	Age int `json:"monster_age"`
	Score float32
	Sex string
	other bool
}

func (s Monster) Print()  {
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end---")
}

func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (s Monster) Set(name string, age int, score float32, sex string)  {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func TestStruct(a interface{})  {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)

	kind := val.Kind()
	if kind != reflect.Struct {
		fmt.Println("参数不是结构体类型")
		return
	}

	num := val.NumField()
	for i := 0; i < num; i++ {
		field := val.Field(i)
		fmt.Printf("field[%d]: %v \n", i, field)

		jsonTag := typ.Field(i).Tag.Get("json")
		if jsonTag != "" {
			fmt.Printf("field[%d] json tag: %v \n", i, jsonTag)
		}
	}

	// 方法的顺序是根据方法名的排序 (依据ASCII码)
	numMethod := val.NumMethod()
	fmt.Println("method num: ", numMethod)

	val.Method(1).Call(nil)

	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(20))

	res := val.Method(0).Call(params)
	fmt.Println(res[0].Int())
}

func main() {
	s := Monster{
		Name:  "jake",
		Age:   34,
		Score: 85,
	}

	TestStruct(s)
}
