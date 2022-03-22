package parser

import (
	"crawler/crawler/engine"
	"regexp"
)

// cityRegexp 城市正则表达式
var cityRegexp = regexp.MustCompile(`<a[^>]+href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`)

// ParseCityList 解析城市列表
func ParseCityList(contents []byte, _ string) engine.ParseResult {
	result := engine.ParseResult{}

	matches := cityRegexp.FindAllSubmatch(contents, -1)
	for _, match := range matches {
		request := engine.Request{
			Url:    string(match[1]),
			Parser: engine.NewFunctionParser(ParseCity, "CityParser"),
		}
		result.Requests = append(result.Requests, request)
	}

	return result
}
