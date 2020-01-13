package parser

import (
	"regexp"
	"studygolang/crawler/engine"
)
                                       // 下一页 href="http://www.zhenai.com/zhenghun/zhenjiang/2"
var cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
var cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
var sexRe = regexp.MustCompile(`<td[^>]*><span class="grayL">性别：</span>([^<]+)</td>`)


func ParseCity(contents []byte) engine.ParseResult {
	result := engine.ParseResult{}

	cityList := cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range cityList {
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	matches := cityRe.FindAllSubmatch(contents, -1)
	sexMatches := sexRe.FindAllSubmatch(contents, -1)
	for i, m := range matches {
		user := m[2]
		sex := sexMatches[i][1]

		result.Items    = append(result.Items, "User " + string(user))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, string(user), string(sex))
			},
		})
	}
	return result
}
