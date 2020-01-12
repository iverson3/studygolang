package parser

import (
	"regexp"
	"strconv"
	"strings"
	"studygolang/crawler/engine"
	"studygolang/crawler/model"
)

// 上海 | 51岁 | 大专 | 离异 | 160cm | 12001-20000元

//const info = `<div data-v-5b109fc3="" class="des f-cl">([^|^<]+) | (\d+)岁 | ([^|^<]+) | ([^|^<]+) | (\d+)cm | ([^|^<]+元)`
const info =`<div class="des f-cl" data-v-3c42fade>([^<]+)</div>`
var infoRe = regexp.MustCompile(info)

func ParseProfile(contents []byte, name string, sex string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	profile.Gender = sex

	match := infoRe.FindSubmatch(contents)
	if match != nil {
		userInfo := string(match[1])
		// 阿坝 | 32岁 | 中专 | 未婚 | 170cm | 50000元以上

		arr := strings.Split(userInfo, "|")
		//fmt.Printf("$%s$", arr[1])
		//profile.Age = len(arr)

		for i := range arr {
			arr[i] = strings.TrimSpace(arr[i])
		}
		arr[1] = strings.TrimRight(arr[1], "岁")
		arr[4] = strings.TrimRight(arr[4], "cm")

		age, err := strconv.Atoi(string(arr[1]))
		if err == nil {
			profile.Age = age
		}

		height, err := strconv.Atoi(string(arr[4]))
		if err == nil {
			profile.Height = height
		}

		profile.Education = string(arr[2])
		profile.Marriage = string(arr[3])
		profile.Income = string(arr[5])
	}

	result := engine.ParseResult{
		Requests: nil,
		Items:    []interface{} {profile},
	}
	return result
}































