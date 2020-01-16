package parser

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"studygolang/crawler/engine"
	"studygolang/crawler/model"
)

// 上海 | 51岁 | 大专 | 离异 | 160cm | 12001-20000元
// `<div data-v-5b109fc3="" class="des f-cl">([^|^<]+) | (\d+)岁 | ([^|^<]+) | ([^|^<]+) | (\d+)cm | ([^|^<]+元)`
var infoRe = regexp.MustCompile(`<div class="des f-cl" data-v-3c42fade>([^<]+)</div>`)
var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func parseProfile(contents []byte, url string, name string, sex string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	profile.Gender = sex

	match := infoRe.FindSubmatch(contents)
	if match != nil {
		userInfo := string(match[1])
		log.Printf("Userinfo: %s", userInfo)

		arr := strings.Split(userInfo, "|")
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
		Items:    []engine.Item{
			{
				Type:    "zhenai",
				Id:      extractString([]byte(url), idUrlRe),
				Url:     url,
				Payload: profile,
			},
		},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}





























