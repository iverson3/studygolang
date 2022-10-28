package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

type FileInfo struct {
	Name string
	Base64Str string
}

func main() {
	var recordId int64 = 1573227028439588864
	phone := "13396095889"
	name := "王帆"
	amount := float64(12750) / 100
	var createTime int64 = 1666799999
	imgBase64Str, err := getImageBase64(recordId, createTime, phone, name, amount)
	if err != nil {
		return
	}

	// 构建post文件信息
	fileInfos := make([]FileInfo, 1)
	fileInfos[0].Name = "shubi.png"
	fileInfos[0].Base64Str = imgBase64Str

	postData := make(map[string]interface{})
	postData["name"] = "stefan"
	postData["files"] = fileInfos
	postBytes, err := json.Marshal(postData)

	// 发起post请求
	req, err := http.NewRequest("post", "http://127.0.0.1:9000/post_file", bytes.NewReader(postBytes))
	req.Header.Set("Content-Type", "application/json")
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
}

func getImageBase64(orderNo, createTime int64, phone, name string, amount float64) (string, error) {
	var templatePng = "https://2g-filebox-daemon-test.oss-cn-shenzhen.aliyuncs.com/vip_park/shubi-template.png"
	var outFile = "./out.png"
	startPointX := 125
	startPointY := 157
	widthDuration := 261
	heightDuration := 42

	// 获取图片文件
	response, err := http.Get(templatePng)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	inBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	infos := make([]*DrawTextInfo, 7)
	infos[0] = &DrawTextInfo{
		Text: strconv.Itoa(int(orderNo)),
		X:    385,
		Y:    71,
	}
	infos[1] = &DrawTextInfo{
		Text: phone,
		X:    startPointX,
		Y:    startPointY,
	}
	infos[2] = &DrawTextInfo{
		Text: name,
		X:    startPointX + widthDuration,
		Y:    startPointY,
	}
	infos[3] = &DrawTextInfo{
		Text: name,
		X:    startPointX,
		Y:    startPointY + heightDuration,
	}
	infos[4] = &DrawTextInfo{
		Text: time.Unix(createTime, 0).Format("2006-01-02 15:04:05"),
		X:    startPointX + widthDuration,
		Y:    startPointY + heightDuration,
	}
	infos[5] = &DrawTextInfo{
		Text: strconv.FormatFloat(amount, 'f', 2, 64),
		X:    startPointX,
		Y:    startPointY + heightDuration * 2,
	}
	infos[6] = &DrawTextInfo{
		Text: phone,
		X:    startPointX + widthDuration,
		Y:    startPointY + heightDuration * 2,
	}
	out, err := DrawStringOnImageAndSave(inBytes, infos)
	if err != nil {
		return "", err
	}

	//保存图片
	fSave, err := os.Create(outFile)
	if err != nil {
		return "", err
	}
	err = jpeg.Encode(fSave, out, nil)
	fSave.Close()
	if err != nil {
		return "", err
	}

	// 读取文件转base64
	imgBytes, err := os.ReadFile(outFile)
	imgBase64Str := base64.StdEncoding.EncodeToString(imgBytes)

	// 删除临时图片文件
	err = os.Remove(outFile)
	if err != nil {
		fmt.Println("remove failed, error:", err)
	}

	return imgBase64Str, nil
}
