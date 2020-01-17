package view

import (
	"os"
	"studygolang/crawler/engine"
	"studygolang/crawler/frontend/model"
	common "studygolang/crawler/model" // 名字冲突了 使用别名
	"testing"
)

func TestTemplate(t *testing.T)  {
	view := CreateSearchResultView("template.html")

	out, err := os.Create("template.test.html")
	if err != nil {
		panic(err)
	}

	item := engine.Item{
		Type:    "zhenai",
		Id:      "1626466343",
		Url:     "http://album.zhenai.com/u/1626466343",
		Payload: common.Profile{
			Name:       "九月",
			Gender:     "女士",
			Age:        37,
			Height:     162,
			Weight:     0,
			Income:     "3000元以下",
			Marriage:   "离异",
			Education:  "中专",
			Occupation: "",
			Hokou:      "",
			Xinzuo:     "",
			House:      "",
			Car:        "",
		},
	}

	page := model.SearchResult{}
	page.Hits = 125
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}
}








































