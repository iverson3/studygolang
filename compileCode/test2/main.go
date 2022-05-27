package main

func Add(a *int64, b int64) int64 {
	var c int64
	c = *a + b
	*a = 8
	b = 9
	return  c
}

func main() {
	var num1, num2 int64
	num1, num2 = 3, 4

	sum := Add(&num1, num2)
	_ = sum
}
