package main

import (
	"context"
	"fmt"
	"reflect"
)

type Person struct {
	Name string
}

type mdOutgoingKey struct {}
type mdIngoingKey struct {}

func main() {
	ctx := context.Background()
	ctx2 := testContext(ctx)

	in := mdIngoingKey{}
	out := mdOutgoingKey{}

	inValue := ctx2.Value(in)
	fmt.Println(inValue)

	outValue := ctx2.Value(out)
	fmt.Println(outValue)

	fmt.Printf("%v \n", in)
	fmt.Println(reflect.TypeOf(in))
	fmt.Printf("%v \n", out)
	fmt.Println(reflect.TypeOf(out))

	return

	p, _ := test("xxx")
	fmt.Println(p.Name)
}

func test(name string) (p Person, err error) {
	fmt.Printf("p: %v \n", p)
	//fmt.Printf("p Name: %v \n", p.Name)
	p.Name = name
	return
}

func testContext(ctx context.Context) context.Context {
	ctx2 := context.WithValue(ctx, mdOutgoingKey{}, "xxx")
	return ctx2
}
