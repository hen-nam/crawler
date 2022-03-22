package parser

import (
	"io/ioutil"
	"testing"
)

// TestParseCityList 测试解析城市列表
func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents, "")

	const resultSize = 470
	expectedUrls := []string{
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/aba",
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/akesu",
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/alashanmeng",
	}

	if size := len(result.Requests); size != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, size)
	}
	for i, expectedUrl := range expectedUrls {
		if url := result.Requests[i].Url; url != expectedUrl {
			t.Errorf("expected url #%d: %s; but was %s", i, expectedUrl, url)
		}
	}
}
