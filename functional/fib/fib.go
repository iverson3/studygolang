package fib

// 定义一个函数类型
type IntGen func() int

// 斐波拉契数列
func Fibonacci() IntGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a + b
		return a
	}
}

// 我的实现
func Fibonacci2() func() int {
	a := 0
	b := 0
	return func() int {
		sum := 1
		if a != 0 && b != 0 {
			sum = a + b
			a, b = b, sum
		}
		if a == 0 {
			a = 1
		} else {
			if b == 0 {
				b = 1
			}
		}

		return sum
	}
}
