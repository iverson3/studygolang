package main

import "testing"

func BenchmarkDoPreAllocMem(b *testing.B) {
	list := createSlice(100)
	repeat := 100

	for i := 0; i < b.N; i++ {
		DoPreAllocMem(repeat, list)
	}
}