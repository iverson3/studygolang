package parser

import (
	"regexp"
	"studygolang/crawler/engine"
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
			ParserFunc: ParseCity,
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
			Url:        string(m[1]),
			ParserFunc: ProfileParser(username, sex),
		})
	}
	return result
}

func ProfileParser(username, sex string) engine.ParserFunc {
	return func(c []byte, url string) engine.ParseResult {
		return ParseProfile(c, url, username, sex)
	}
}























