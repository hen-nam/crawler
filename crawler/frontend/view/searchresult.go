package view

import (
	"crawler/crawler/frontend/model"
	"html/template"
	"io"
)

// SearchResultView 搜索结果视图
type SearchResultView struct {
	template *template.Template
}

// Render 渲染模板
func (view SearchResultView) Render(writer io.Writer, data model.SearchResult) error {
	return view.template.Execute(writer, data)
}

// CreateSearchResultView 创建搜索结果视图
func CreateSearchResultView(filename string) SearchResultView {
	return SearchResultView{
		template: template.Must(template.ParseFiles(filename)),
	}
}
