package parser

import (
	"regexp"
	"studygolang/crawler/engine"
	"studygolang/crawler_distributed/config"
)
                                       // 下一页 href="http://www.zhenai.com/zhenghun/zhenjiang/2"
var cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
var cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
var sexRe = regexp.MustCompile(`<td[^>]*><span class="grayL">性别：</span>([^<]+)</td>`)


func ParseCity(contents []byte, url string) engine.ParseResult {
	result := engine.ParseResult{}

	// 匹配用户列表页面中的"下一页"和"其他类型的用户列表"对应的链接url
	cityList := cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range cityList {
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
	}

	// 匹配用户昵称和用户详情页的url 以及用户性别
	matches := cityRe.FindAllSubmatch(contents, -1)
	sexMatches := sexRe.FindAllSubmatch(contents, -1)
	for i, m := range matches {
		username := string(m[2])
		sex := string(sexMatches[i][1])

		//result.Items    = append(result.Items, "User " + string(user))
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: NewProfileParser(username, sex),
	})
	}
	return result
}


type ProfileParser struct {
	UserName string
	Sex string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return parseProfile(contents, url, ProfileParser{
		UserName: p.UserName,
		Sex:      p.Sex,
	})
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return config.ParseProfile, ProfileParser{
		UserName: p.UserName,
		Sex:      p.Sex,
	}
}

func NewProfileParser(userName, sex string) *ProfileParser {
	return &ProfileParser{
		UserName: userName,
		Sex:      sex,
	}
}























