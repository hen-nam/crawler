package parser

import (
	"crawler/crawler/engine"
	"regexp"
)

var (
	// profileRegexp 用户正则表达式
	profileRegexp = regexp.MustCompile(`<a[^>]+href="(http://localhost:8080/mock/album.zhenai.com/u/\d+)"[^>]*>([^<]+)</a>`)
	// otherCityRegexp 城市正则表达式
	otherCityRegexp = regexp.MustCompile(`<a[^>]+href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[^"]+)"[^>]*>[^<]+</a>`)
)

// ParseCity 解析城市
func ParseCity(contents []byte, _ string) engine.ParseResult {
	result := engine.ParseResult{}

	matches := profileRegexp.FindAllSubmatch(contents, -1)
	for _, match := range matches {
		request := engine.Request{
			Url:    string(match[1]),
			Parser: NewProfileParser(string(match[2])),
		}
		result.Requests = append(result.Requests, request)
	}

	matches = otherCityRegexp.FindAllSubmatch(contents, -1)
	for _, match := range matches {
		request := engine.Request{
			Url:    string(match[1]),
			Parser: engine.NewFunctionParser(ParseCity, "CityParser"),
		}
		result.Requests = append(result.Requests, request)
	}

	return result
}
