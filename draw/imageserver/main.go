package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type PostData struct {
	Name string `json:"name"`
	FileList []FileInfo `json:"files"`
}

type FileInfo struct {
	Name string
	Base64Str string
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/post_file", postFile)

	server := &http.Server{
		Addr:              ":9000",
		Handler:           mux,
		ReadTimeout:       5,
		WriteTimeout:      5,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func postFile(w http.ResponseWriter, r *http.Request) {

	all, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	data := new(PostData)
	err = json.Unmarshal(all, data)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Println(data.Name)

	fSave, err := os.OpenFile(data.FileList[0].Name, os.O_CREATE | os.O_RDWR | os.O_TRUNC, os.ModePerm)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	defer fSave.Close()

	imgBytes, err := base64.StdEncoding.DecodeString(data.FileList[0].Base64Str)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	_, err = fSave.Write(imgBytes)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("done"))
}