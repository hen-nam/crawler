package controller

import (
	"context"
	"crawler/crawler/engine"
	"crawler/crawler/frontend/model"
	"crawler/crawler/frontend/view"
	"github.com/olivere/elastic/v7"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// SearchResultHandler 搜索结果处理器
type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

// ServeHTTP 启动 HTTP 服务
func (handler SearchResultHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	query := strings.TrimSpace(request.FormValue("query"))
	from, err := strconv.Atoi(request.FormValue("from"))
	if err != nil {
		from = 0
	}

	data, err := handler.getSearchResult(query, from)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	err = handler.view.Render(writer, data)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
}

// getSearchResult 获取模型数据
func (handler SearchResultHandler) getSearchResult(query string, from int) (model.SearchResult, error) {
	data := model.SearchResult{
		Query: query,
	}

	query = rewriteQuery(query)
	q := elastic.NewQueryStringQuery(query)
	service := handler.client.Search("zhenai").Query(q).From(from)
	result, err := service.Do(context.Background())
	if err != nil {
		return data, err
	}

	data.From = from
	const fromOffset = 10
	data.PrevFrom = from - fromOffset
	data.NextFrom = from + fromOffset
	data.Hits = int(result.TotalHits())
	items := result.Each(reflect.TypeOf(engine.Item{}))
	for _, item := range items {
		data.Items = append(data.Items, item.(engine.Item))
	}
	return data, nil
}

// CreateSearchResultHandler 创建搜索结果处理器
func CreateSearchResultHandler(filename string) SearchResultHandler {
	view := view.CreateSearchResultView(filename)

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	return SearchResultHandler{
		view:   view,
		client: client,
	}
}

// rewriteQuery 重写查询条件
func rewriteQuery(query string) string {
	regexp := regexp.MustCompile(`([A-z][a-z]*):`)
	return regexp.ReplaceAllString(query, "Payload.$1:")
}
