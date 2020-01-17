package parser

import (
	"fmt"
	"io/ioutil"
	"studygolang/crawler/engine"
	"studygolang/crawler/model"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := parseProfile(contents, "http://album.zhenai.com/u/1626466343", ProfileParser{
		UserName: "九月",
		Sex: "女士",
	})
	
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result.Items)
	}
	actual := result.Items[0]

	fmt.Printf("%v", result)

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

	if actual != expected {
		t.Errorf("expected %v; but was %v", expected, actual)
	}
}


























