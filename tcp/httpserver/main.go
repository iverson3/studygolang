package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/index", indexFunc)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func indexFunc(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	var resp string
	if name == "" {
		resp = "response from server: invalid params"
	} else {
		resp = fmt.Sprintf("response from server: %s", name)
	}
	_, _ = w.Write([]byte(resp))
}

