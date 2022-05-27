package pkg1

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
	g = c - d
	var h int64 = 10
	return g, h
}