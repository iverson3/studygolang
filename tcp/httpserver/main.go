package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/go-yx-extension/api/v1/test/test/getname", indexFunc)

	err := http.ListenAndServe(":9099", nil)
	if err != nil {
		panic(err)
	}
}

func indexFunc(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	fmt.Printf("got a request: %s\n", name)
	time.Sleep(3 * time.Second)

	var resp string
	if name == "" {
		resp = "response from server: invalid params"
	} else {
		resp = fmt.Sprintf("response from server: %s", name)
	}
	_, _ = w.Write([]byte(resp))
}

