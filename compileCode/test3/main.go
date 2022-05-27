package main

func Add(a int64) int64 {
	var b int64

	defer func() {
		a++
		b++
	}()

	a++
	b = a
	return b
}

func main() {
	var num1, num2 int64
	num2 = Add(num1)

	_ = num2
}
