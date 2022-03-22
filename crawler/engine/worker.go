package engine

import (
	"crawler/crawler/fetcher"
	"log"
)

// Worker 工作器
func Worker(request Request) (ParseResult, error) {
	log.Printf("Fetching url: %s", request.Url)
	contents, err := fetcher.Fetch(request.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", request.Url, err)
		return ParseResult{}, err
	}

	result := request.Parser.Parse(contents, request.Url)
	return result, nil
}
