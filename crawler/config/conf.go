package config

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"runtime"
)

type configuration struct {
	ElasticServerUrl string
	SeedUrl    string
	ElasticSearchIndex string
	RpcServeHost string
}

// 目前弃用
// 如果配置文件读取失败或配置文件解析失败 则直接panic
func Config() configuration {
	currentPath := getCurrentPath()
	// 如果有多个配置文件 则把这个变量做为参数
	//configFilename := "crawler/config/server.json"
	configFilename := "server.json"
	file, err := os.Open(currentPath + "/" + configFilename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := configuration{}
	err = decoder.Decode(&conf)
	if err != nil {
		log.Printf("json decode error: %v", err)
		panic(err)
	}
	return conf
}

// 获取当前文件的完整路径
func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}