package parser

import (
	"crawler/crawler/engine"
	"crawler/crawler/model"
	"io/ioutil"
	"testing"
)

func TestProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := parseProfile(contents, "http://localhost:8080/mock/album.zhenai.com/u/8256018539338750764", "寂寞成影萌宝")

	expectedItem := engine.Item{
		Url:   "http://localhost:8080/mock/album.zhenai.com/u/8256018539338750764",
		Index: "zhenai",
		Id:    "8256018539338750764",
		Payload: model.Profile{
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

	if size := len(result.Items); size != 1 {
		t.Errorf("Items should contain 1 element; but was %d", size)
	}
	actualItem := result.Items[0]
	if actualItem != expectedItem {
		t.Errorf("expected Item was %+v; but was %+v", expectedItem, actualItem)
	}
}
