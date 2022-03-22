package parser

import (
	"crawler/crawler/engine"
	"crawler/crawler/model"
	"crawler/crawler_distributed/config"
	"regexp"
	"strconv"
)

var (
	genderRegexp        = regexp.MustCompile(`<span[^>]*>性别[^>]*</span><span[^>]*>([^<]+)</span>`)
	ageRegexp           = regexp.MustCompile(`<span[^>]*>年龄[^>]*</span>(\d+)岁`)
	heightRegexp        = regexp.MustCompile(`<span[^>]*>身高[^>]*</span><span[^>]*>(\d+)CM</span>`)
	weightRegexp        = regexp.MustCompile(`<span[^>]*>体重[^>]*</span><span[^>]*>(\d+)KG</span>`)
	incomeRegexp        = regexp.MustCompile(`<span[^>]*>月收入[^>]*</span>([^<]+)`)
	marriageRegexp      = regexp.MustCompile(`<span[^>]*>婚况[^>]*</span>([^<]+)`)
	educationRegexp     = regexp.MustCompile(`<span[^>]*>学历[^>]*</span>([^<]+)`)
	occupationRegexp    = regexp.MustCompile(`<span[^>]*>职业[^>]*</span>([^<]+)`)
	locationRegexp      = regexp.MustCompile(`<span[^>]*>籍贯[^>]*</span>([^<]+)`)
	constellationRegexp = regexp.MustCompile(`<span[^>]*>星座[^>]*</span><span[^>]*>([^<]+)</span>`)
	houseRegexp         = regexp.MustCompile(`<span[^>]*>住房条件[^>]*</span><span[^>]*>([^<]+)</span>`)
	carRegexp           = regexp.MustCompile(`<span[^>]*>是否购车[^>]*</span><span[^>]*>([^<]+)</span>`)
	otherProfileRegexp  = regexp.MustCompile(`<a[^>]+href="(http://localhost:8080/mock/album.zhenai.com/u/\d+)"[^>]*>([^<]+)</a>`)
	idRegexp            = regexp.MustCompile(`http://localhost:8080/mock/album.zhenai.com/u/(\d+)`)
)

// parseProfile 解析用户
func parseProfile(contents []byte, url string, name string) engine.ParseResult {
	result := engine.ParseResult{}

	item := engine.Item{
		Url:   url,
		Index: "zhenai",
		Id:    extractString([]byte(url), idRegexp),
		Payload: model.Profile{
			Name:          name,
			Gender:        extractString(contents, genderRegexp),
			Age:           extractInt(contents, ageRegexp),
			Height:        extractInt(contents, heightRegexp),
			Weight:        extractInt(contents, weightRegexp),
			Income:        extractString(contents, incomeRegexp),
			Marriage:      extractString(contents, marriageRegexp),
			Education:     extractString(contents, educationRegexp),
			Occupation:    extractString(contents, occupationRegexp),
			Location:      extractString(contents, locationRegexp),
			Constellation: extractString(contents, constellationRegexp),
			House:         extractString(contents, houseRegexp),
			Car:           extractString(contents, carRegexp),
		},
	}
	result.Items = append(result.Items, item)

	matches := otherProfileRegexp.FindAllSubmatch(contents, -1)
	for _, match := range matches {
		request := engine.Request{
			Url:    string(match[1]),
			Parser: NewProfileParser(string(match[2])),
		}
		result.Requests = append(result.Requests, request)
	}

	return result
}

// extractString 提取字符串
func extractString(contents []byte, r *regexp.Regexp) string {
	match := r.FindSubmatch(contents)
	if match == nil || len(match) < 2 {
		return ""
	}
	return string(match[1])
}

// extractInt 提取整数
func extractInt(contents []byte, r *regexp.Regexp) int {
	s := extractString(contents, r)
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

// ProfileParser 用户解析器
type ProfileParser struct {
	name string
}

// Parse 解析
func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return parseProfile(contents, url, p.name)
}

// Serialize 序列化
func (p *ProfileParser) Serialize() (string, interface{}) {
	return config.ProfileParser, p.name
}

// NewProfileParser 创建用户解析器
func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		name: name,
	}
}
