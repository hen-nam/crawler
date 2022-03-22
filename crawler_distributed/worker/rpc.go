package worker

import "crawler/crawler/engine"

// WorkerService 工作器服务
type WorkerService struct{}

func (WorkerService) Process(request Request, result *ParseResult) error {
	req, err := DeserializeRequest(request)
	if err != nil {
		return err
	}

	res, err := engine.Worker(req)
	if err != nil {
		return err
	}

	*result = SerializeParseResult(res)
	return nil
}
