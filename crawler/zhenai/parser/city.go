package parser

import (
	"regexp"
	"studygolang/crawler/engine"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

const sexRe = `<td[^>]*><span class="grayL">性别：</span>([^<]+)</td>`

func ParseCity(contents []byte) engine.ParseResult {
	compile := regexp.MustCompile(cityRe)
	matches := compile.FindAllSubmatch(contents, -1)

	sexCompile := regexp.MustCompile(sexRe)
	sexMatches := sexCompile.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
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
