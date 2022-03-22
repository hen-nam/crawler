package model

import "crawler/crawler/engine"

// SearchResult 搜索结果
type SearchResult struct {
	Query    string
	From     int
	PrevFrom int
	NextFrom int
	Hits     int
	Items    []engine.Item
}
