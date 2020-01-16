package main

import (
	"net/http"
	"studygolang/crawler/frontend/controller"
)

func main() {
	// 处理静态资源文件 (css js 图片 )
	// 比如 http://localhost:8888/css/bootstrap.min.css   http://localhost:8888/js/bootstrap.min.js
	http.Handle("/", http.FileServer(
		http.Dir("crawler/frontend/view")))
	// 处理动态搜索请求
	http.Handle("/search",
		controller.CreateSearchResultHandler("crawler/frontend/view/template.html"))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
















































