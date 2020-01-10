package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func variableShorter()  {
	a, b, c, s := 1, 3, true, "string"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler()  {
	fmt.Printf("%.3f\n", cmplx.Exp(1i * math.Pi) + 1)
}

func triangle()  {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a * a) + float64(b * b)))
	return c
}

func main() {
	variableShorter()
	euler()

	triangle()
}
