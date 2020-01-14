package persist

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	"studygolang/crawler/model"
	"testing"
)

func TestSaver(t *testing.T) {
	expected := model.Profile{
		Name:       "测试用户",
		Gender:     "男性",
		Age:        25,
		Height:     175,
		Weight:     62,
		Income:     "10000-15000",
		Marriage:   "未婚",
		Education:  "本科",
		Occupation: "",
		Hokou:      "",
		Xinzuo:     "",
		House:      "",
		Car:        "",
	}
	id, err := save(expected)
	if err != nil {
		panic(err)
	}

	// TODO: Try to start up elastic search
	// here using docker go client.
	client, err := elastic.NewClient(
		elastic.SetURL("http://47.107.149.234:9200"),
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index("dating_profile"). // 数据库
		Type("zhenai"). // 表名
		Id(id). // 数据 (不设置id  让系统自动生成)
		Do(context.Background()) // 后台运行

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", resp)
	fmt.Printf("%s", resp.Source)

	var actual model.Profile
	err = json.Unmarshal(resp.Source, &actual)
	if err != nil {
		panic(err)
	}

	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}



































