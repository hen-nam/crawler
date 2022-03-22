package main

import (
	"crawler/crawler/engine"
	"crawler/crawler/persist"
	"crawler/crawler/scheduler"
	"crawler/crawler/zhenai/parser"
)

// main 执行
func main() {
	itemChan, err := persist.ItemSaver()
	if err != nil {
		panic(err)
	}

	//e := engine.SimpleEngine{}
	e := engine.ConcurrentEngine{
		//Scheduler: &scheduler.SimpleScheduler{},
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		Processor:   engine.Worker,
		ItemChan:    itemChan,
	}
	request := engine.Request{
		Url:    "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		Parser: engine.NewFunctionParser(parser.ParseCityList, "CityListParser"),
		//Url:    "http://localhost:8080/mock/www.zhenai.com/zhenghun/aba",
		//Parser: engine.NewFunctionParser(parser.ParseCity, "CityParser"),
		//Url:    "http://localhost:8080/mock/album.zhenai.com/u/8256018539338750764",
		//Parser: parser.NewProfileParser("寂寞成影萌宝"),
	}
	e.Run(request)
}
