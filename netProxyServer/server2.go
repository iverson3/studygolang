package main

import "net/http"

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("web222"))
	})
	http.ListenAndServe(":8882", nil)
}
