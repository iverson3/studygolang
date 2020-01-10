package mock

import "fmt"

type Retriever struct {
	Contents string
}

// 实现系统的stringer接口的String方法
func (r *Retriever) String() string {
	return fmt.Sprintf("Retriever: {Contents=%s}", r.Contents)
}

// 实现Poster接口的Post方法
func (r *Retriever) Post(url string, form map[string]string) string {
	// 要实现下面的修改对调用者起作用 必须使用 *Retriever引用传递 (否则就是值传递 修改不会影响到外面)
	r.Contents = form["contents"]
	return "ok"
}

// 实现Retriever接口的Get方法
// 在go语言中 只要实现了接口的方法 即表示实现了该接口 (没有太多的限制和约束)
func (r *Retriever) Get(url string) string {
	return r.Contents
}

