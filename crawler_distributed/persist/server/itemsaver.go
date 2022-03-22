package main

import (
	"crawler/crawler_distributed/persist"
	"crawler/crawler_distributed/rpcsupport"
	"flag"
	"fmt"
	"github.com/olivere/elastic/v7"
)

var port = flag.Int("port", 0, "the port to listen on")

// main 启动服务
// go run crawler_distributed/persist/server/itemsaver.go --port=1234
func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}

	address := fmt.Sprintf(":%d", *port)
	err := serveRpc(address)
	if err != nil {
		panic(err)
	}
}

// serveRpc 启动服务
func serveRpc(address string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}

	service := persist.ItemSaverService{
		Client: client,
	}
	err = rpcsupport.ServeRpc(address, &service)
	return err
}
