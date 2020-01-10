package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"studygolang/errhandling/filelistingserver/filelisting"
)

// 文件列表服务器
// 输出指定文件的内容

// 定义userError接口
type userError interface {
	error
	Message() string
}

//func errUserError(writer http.ResponseWriter, request *http.Request) error {
//	return testingUserError("user error")
//}

//func TestErrWrapper(t *testing.T) {
//	tests := []struct{
//		h appHandler
//		code int
//		message string
//	} {
//		{errPanic, 500, "Internal Server Error"},
//		{errUserError, 400, "user error"},
//	}
//}

type appHandler func(writer http.ResponseWriter, request *http.Request) error

// errWrapper() 体现了典型的函数式编程，即参数是函数 返回值也是函数
// 统一的错误处理函数
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		// recover可能出现的错误panic
		defer func() {
			r := recover()
			if r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		if err != nil {
			log.Printf("Error handling request: %s", err.Error())

			// 将handler函数返回的error转userError 如果成功则表示是一个userError 则进行相应的处理
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			// 判断不同类型的错误
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			case os.IsTimeout(err):
				code = http.StatusRequestTimeout
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	// 第一个参数是服务器的根目录 它指向awesomeProject1目录下
	// 浏览器访问时url http://localhost:8888/list/fib.txt
	http.HandleFunc("/list/", errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}



























