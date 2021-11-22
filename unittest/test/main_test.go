package main

import "testing"

func TestFoo(t *testing.T) {
	foo := Foo(1, 2)
	wanted := 3
	if foo != wanted {
		t.Fatalf("call function Foo(); want %d, but got %d\n", wanted, foo)
	}
}

func BenchmarkFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Foo(2, 3)
	}
}

func TestTestQuickSort(t *testing.T) {
	source := []int{1, 4, 8, 3, 5, 8, 6, 1, 2}
	res := TestQuickSort(source)
	wanted := []int{1, 1, 2, 3, 4, 5, 6, 8, 8}

	res1 := len(wanted) == len(res)
	if !res1 {
		t.Fatalf("call111 function TestQuickSort; wanted: %v, but got %v\n", wanted, res)
	}

	res2 := true
	for i := 0; i < len(wanted); i++ {
		if wanted[i] != res[i] {
			res2 = false
			break
		}
	}
	if !res2 {
		t.Fatalf("call222 function TestQuickSort; wanted: %v, but got %v\n", wanted, res)
	}
}
