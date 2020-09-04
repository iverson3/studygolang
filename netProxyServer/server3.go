package main

import "net/http"

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("web333"))
	})
	http.ListenAndServe(":8883", nil)
}
