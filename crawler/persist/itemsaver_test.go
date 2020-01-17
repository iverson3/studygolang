package persist

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/olivere/elastic.v6"
	"studygolang/crawler/engine"
	"studygolang/crawler/model"
	"studygolang/crawler_distributed/config"
	"testing"
)

func TestSaver(t *testing.T) {
	expected := engine.Item{
		Type:    "zhenai",
		Id:      "1626466343",
		Url:     "http://album.zhenai.com/u/1626466343",
		Payload: model.Profile{
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

	// TODO: Try to start up elastic search
	// here using docker go client.
	client, err := elastic.NewClient(
		elastic.SetURL(config.ElasticServerUrl),
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	err = Save(client, expected, index)
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index(index). // 数据库
		Type(expected.Type). // 表名
		Id(expected.Id). // 数据 (不设置id  让系统自动生成)
		Do(context.Background()) // 后台运行

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", resp)
	fmt.Printf("%s", resp.Source)

	var actual engine.Item
	json.Unmarshal(*resp.Source, &actual)
	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}



































