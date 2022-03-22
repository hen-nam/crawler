package main

import (
	"crawler/crawler/engine"
	"crawler/crawler/model"
	"crawler/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

// TestItemSaver 测试项目存储服务
func TestItemSaverService(t *testing.T) {
	const address = ":1234"

	go serveRpc(address)
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(address)
	if err != nil {
		panic(err)
	}

	item := engine.Item{
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
	result := ""
	err = client.Call("ItemSaverService.Save", item, &result)
	if err != nil {
		t.Errorf("result: %s, error: %v", result, err)
	}
}
