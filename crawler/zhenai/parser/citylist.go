package parser

import (
	"regexp"
	"studygolang/crawler/engine"
	"studygolang/crawler_distributed/config"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, url string) engine.ParseResult {
	compile := regexp.MustCompile(cityListRe)
	matches := compile.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	limit := 500
	for _, m := range matches {
		//result.Items    = append(result.Items, "City " + string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
		limit--
		if limit == 0 {
			break
		}
	}
	return result
}




























