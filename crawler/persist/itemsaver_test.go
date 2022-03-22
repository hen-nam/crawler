package persist

import (
	"context"
	"crawler/crawler/engine"
	"crawler/crawler/model"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"testing"
)

func TestSave(t *testing.T) {
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

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	err = Save(client, expectedItem)
	if err != nil {
		panic(err)
	}

	service := client.Get().Index(expectedItem.Index).Id(expectedItem.Id)
	result, err := service.Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%s", result.Source)

	actualItem := engine.Item{}
	err = json.Unmarshal(result.Source, &actualItem)
	if err != nil {
		panic(err)
	}

	actualItem.Payload, err = model.FromJsonObject(actualItem.Payload)
	if err != nil {
		panic(err)
	}

	if actualItem != expectedItem {
		t.Errorf("got %+v, expected %+v", actualItem, expectedItem)
	}
}
