package utils

import (
	"fmt"
	"github.com/go-ini/ini"
)

var ProxyConfigs map[string]PassInfo

type PassInfo struct {
	Url string  // server地址
	Weight int  // 权重
}

// 读取.ini配置文件
// 将配置项存放到map里面 map[key]value
func init() {
	ProxyConfigs = make(map[string]PassInfo)

	EnvConfig, err := ini.Load("./utils/env.ini")
	if err != nil {
		fmt.Println("load ini failed! error: ", err)
		return
	}

	proxy, err := EnvConfig.GetSection("proxy") // 获取分区
	if err != nil {
		fmt.Println("get 'proxy' section failed! error: ", err)
		return
	}

	sections := proxy.ChildSections() // 获取proxy分区下的子分区
	for _, sec := range sections {
		path, _   := sec.GetKey("path")   // 获取path配置项的值
		pass, _   := sec.GetKey("pass")   // 获取pass配置项的值
		weight, _ := sec.GetKey("weight") // 获取weight配置项的值
		if path != nil && pass != nil {
			passInfo := PassInfo{
				Url:    pass.Value(),
				Weight: 1,  // 默认权重为 1
			}
			if weight != nil {
				i, _ := weight.Int()
				passInfo.Weight = i
			}
			ProxyConfigs[path.Value()] = passInfo
		}
	}
}