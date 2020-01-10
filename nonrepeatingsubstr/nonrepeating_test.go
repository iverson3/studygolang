package main

import "testing"

// 表格驱动测试(Test)
// 代码覆盖率(Coverage)
func TestSubStr(t *testing.T)  {
	tests := []struct {
		s string
		ans int
	} {
		// Normal case
		{ "aaaaagfa", 3},
		{ "abcabcabc", 3},

		// Edge case
		{"", 0},
		{"b", 1},
		{"aaaaaaa", 1},
		{"abcabcabcd", 4},
		{"abcdefgh", 8},

		// Chinese support
		{"一二三二一", 3},
		{"这里是北京", 5},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花",8},
	}

	for _, tt := range tests {
		if actual := lengthOfNonRepeatingSubStr(tt.s); actual != tt.ans {
			t.Errorf("got %d for input %s; expected %d", actual, tt.s, tt.ans)
		}
	}
}

// 性能测试 (Bench)
func BenchmarkSubStr(b *testing.B)  {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	for i := 0; i < 13; i++ {
		s = s + s
	}
	b.Logf("len(s) = %d", len(s))
	ans := 8

	b.ResetTimer()
	// 循环运行多少次由系统决定
	for i := 0; i < b.N; i++ {
		if actual := lengthOfNonRepeatingSubStr(s); actual != ans {
			b.Errorf("got %d for input %s; expected %d", actual, s, ans)
		}
	}
}


















