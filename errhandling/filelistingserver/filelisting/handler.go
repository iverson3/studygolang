package filelisting

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	index := strings.Index(request.URL.Path, prefix)
	log.Println(index)
	if index != 0 {
		return userError("path must start with " + prefix)
	}
	// request.URL.Path is /list/fib.txt
	// 获取文件本身的path ( fib.txt )
	path := request.URL.Path[len(prefix):]

	file, err := os.Open(path)
	if err != nil {
		//http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	// 响应文件内容
	writer.Write(all)
	return nil
}
