package main

import (
	"crawler/crawler_distributed/rpcsupport"
	"crawler/crawler_distributed/worker"
	"flag"
	"fmt"
)

var port = flag.Int("port", 0, "the port to listen on")

// main 启动服务
// go run crawler_distributed/worker/server/worker.go --port=9000
func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}

	address := fmt.Sprintf(":%d", *port)
	err := rpcsupport.ServeRpc(address, worker.WorkerService{})
	if err != nil {
		panic(err)
	}
}
