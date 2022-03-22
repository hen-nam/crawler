package main

import (
	"crawler/crawler_distributed/config"
	"crawler/crawler_distributed/rpcsupport"
	"crawler/crawler_distributed/worker"
	"fmt"
	"testing"
	"time"
)

// TestWorkerService 测试工作器服务
func TestWorkerService(t *testing.T) {
	address := ":9000"

	go rpcsupport.ServeRpc(address, worker.WorkerService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(address)
	if err != nil {
		panic(err)
	}

	request := worker.Request{
		Url: "http://localhost:8080/mock/album.zhenai.com/u/8256018539338750764",
		Parser: worker.SerializedParser{
			Name: config.ProfileParser,
			Args: "寂寞成影萌宝",
		},
	}
	result := worker.ParseResult{}
	err = client.Call(config.WorkerMethod, request, &result)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v\n", result)
}
