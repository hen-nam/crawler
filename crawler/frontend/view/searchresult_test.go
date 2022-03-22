package view

import (
	"crawler/crawler/engine"
	"crawler/crawler/frontend/model"
	common "crawler/crawler/model"
	"os"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView("template.html")

	file, err := os.Create("template.test.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := model.SearchResult{
		Hits: 123,
	}
	item := engine.Item{
		Url:   "http://localhost:8080/mock/album.zhenai.com/u/8256018539338750764",
		Index: "zhenai",
		Id:    "8256018539338750764",
		Payload: common.Profile{
			Name:          "寂寞成影萌宝",
			Gender:        "女",
			Age:           83,
			Height:        105,
			Weight:        137,
			Income:        "财务自由",
			Marriage:      "离异",
			Education:     "初中",
			Occupation:    "金融",
			Location:      "南京市",
			Constellation: "狮子座",
			House:         "无房",
			Car:           "无车",
		},
	}
	for i := 0; i < 10; i++ {
		data.Items = append(data.Items, item)
	}

	err = view.Render(file, data)
	if err != nil {
		t.Error(err)
	}
}
