package controller

import (
	"context"
	"gopkg.in/olivere/elastic.v6"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"studygolang/crawler/engine"
	"studygolang/crawler/frontend/model"
	view2 "studygolang/crawler/frontend/view"
	"studygolang/crawler_distributed/config"
)

type SearchResultHandler struct {
	view view2.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(
		elastic.SetURL(config.ElasticServerUrl),
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view2.CreateSearchResultView(template),
		client: client,
	}
}

// localhost:8888/search?q=男 未婚 博士&from=20
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}

	//fmt.Fprintf(w, "q=%s, from=%d", q, from)

	var page model.SearchResult
	page, err = h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult

	resp, err := h.client.
		Search("dating_profile").
		Query(elastic.NewQueryStringQuery(rewriteQueryString(q))).
		From(from).
		Do(context.Background())

	if err != nil {
		return result, err
	}

	result.Hits = int(resp.TotalHits())
	result.Start = from
	result.Query = q

	for _, v := range resp.Each(reflect.TypeOf(engine.Item{})) {
		item := v.(engine.Item)
		result.Items = append(result.Items, item)
	}
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)

	return result, nil
}

// 将请求参数字符串中的 Age:(<23) 替换为 Payload.Age:(<23)
func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")
}











































