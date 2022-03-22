package main

import (
	"crawler/crawler/engine"
	"crawler/crawler/scheduler"
	"crawler/crawler/zhenai/parser"
	persist "crawler/crawler_distributed/persist/client"
	"crawler/crawler_distributed/rpcsupport"
	worker "crawler/crawler_distributed/worker/client"
	"flag"
	"log"
	"net/rpc"
	"strings"
)

var (
	workerAddresses  = flag.String("worker_addresses", "", "worker addresses (comma separated)")
	itemSaverAddress = flag.String("itemsaver_address", "", "itemsaver address")
)

// main 执行
// go run crawler_distributed/main.go --worker_addresses=:9000,:9001,:9002 --itemsaver_address=:1234
func main() {
	flag.Parse()

	addresses := strings.Split(*workerAddresses, ",")
	pool := createClientPool(addresses)
	processor := worker.CreateProcessor(pool)

	itemChan, err := persist.ItemSaver(*itemSaverAddress)
	if err != nil {
		panic(err)
	}

	//e := engine.SimpleEngine{}
	e := engine.ConcurrentEngine{
		//Scheduler: &scheduler.SimpleScheduler{},
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		Processor:   processor,
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

// createClientPool 创建客户端池
func createClientPool(addresses []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, address := range addresses {
		client, err := rpcsupport.NewClient(address)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", address)
		} else {
			log.Printf("Error connecting to %s: %v", address, err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
