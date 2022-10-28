package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

// 镜像仓库模拟客户端

func main() {
	err := uploadFile()
	fmt.Println(err)

	//err := pullFile()
	//fmt.Println(err)
}

func pullFile() (err error) {
	url := "http://127.0.0.1:8888/images/pull"

	fileName := "xxx.ttf"
	tag := "1.2.0"

	postData := make(map[string]interface{})
	postData["filename"] = fileName
	postData["tag"] = tag
	postBytes, err := json.Marshal(postData)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(postBytes))
	if err != nil {
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// 文件的数据长度
	length := resp.ContentLength

	fileBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if len(fileBytes) != int(length) {
		fmt.Println(len(fileBytes), length)
		return errors.New("file length is wrong")
	}

	dstFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		return
	}
	defer dstFile.Close()

	n, err := dstFile.Write(fileBytes)
	if err != nil {
		return
	}
	if n != len(fileBytes) {
		return errors.New("write file length is wrong")
	}

	fmt.Println("download success")
	return nil
}

func uploadFile() (err error) {
	url := "http://127.0.0.1:8888/images/push"

	imageFile, err := os.Open("./simfang.ttf")
	if err != nil {
		panic(err)
	}

	// 模拟客户端提交表单
	values := map[string]io.Reader {
		"file": imageFile,
		"filename": strings.NewReader("xxx.ttf"),
		"tag": strings.NewReader("1.2.0"),
	}

	// 构建multipart，然后post给服务端
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)
	for key, reader := range values {
		var fw io.Writer
		if x, ok := reader.(io.Closer); ok {
			defer x.Close()
		}

		if x, ok := reader.(*os.File); ok {
			// 添加文件
			if fw, err = writer.CreateFormFile(key, x.Name()); err != nil {
				return
			}
		} else {
			// 添加字符串
			if fw, err = writer.CreateFormField(key); err != nil {
				return
			}
		}

		if _, err = io.Copy(fw, reader); err != nil {
			return
		}
	}
	// close动作会在末端写入 boundary
	writer.Close()

	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return
	}

	// form-data格式，自动生成分隔符
	// 例如 Content-Type: multipart/form-data; boundary=d76.....d29
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		return
	}

	fmt.Println("ok")
	return nil
}