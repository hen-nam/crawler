package client

import (
	"crawler/crawler/engine"
	"crawler/crawler_distributed/config"
	"crawler/crawler_distributed/worker"
	"net/rpc"
)

// CreateProcessor 创建请求处理器
func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	processor := func(request engine.Request) (engine.ParseResult, error) {
		client := <-clientChan
		req := worker.SerializeRequest(request)
		res := worker.ParseResult{}
		err := client.Call(config.WorkerMethod, req, &res)
		if err != nil {
			return engine.ParseResult{}, err
		}

		result := worker.DeserializeParseResult(res)
		return result, nil
	}
	return processor
}
