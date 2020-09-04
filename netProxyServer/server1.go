package main

import "net/http"

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("web111"))
	})
	http.ListenAndServe(":8881", nil)
}
