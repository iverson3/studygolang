package main

//go:noinline
func Add(a, b int64) (f ,i int64) {
	var e int64
	e = a + b
	f ,i = Del(e, b)
	return
}

//go:noinline
func Del(c, d int64) (int64, int64) {
	var g int64
	if c > d {
		g = c - d
	} else {
		g = d - c
	}
	var h int64 = 10
	return g, h
}

func main() {
	var num1 int64 = 3
	var num2 int64 = 4
	_, _ = Add(num1, num2)

	go func() {
		num3 := num1
		_ = num3
	}()
}
