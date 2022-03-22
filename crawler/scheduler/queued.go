package scheduler

import "crawler/crawler/engine"

// QueuedScheduler 队列调度器
type QueuedScheduler struct {
	in      chan engine.Request
	outChan chan chan engine.Request
}

// Submit 提交请求
func (s *QueuedScheduler) Submit(request engine.Request) {
	s.in <- request
}

// WorkerChan 获取连接工作器的通道
func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	out := make(chan engine.Request)
	return out
}

// WorkerReady 准备连接工作器的通道
func (s *QueuedScheduler) WorkerReady(out chan engine.Request) {
	s.outChan <- out
}

// Run 运行
func (s *QueuedScheduler) Run() {
	s.in = make(chan engine.Request)
	s.outChan = make(chan chan engine.Request)

	go func() {
		var requestQueue []engine.Request
		var outQueue []chan engine.Request

		for {
			var activeRequest engine.Request
			var activeOut chan engine.Request
			if len(requestQueue) > 0 && len(outQueue) > 0 {
				activeRequest = requestQueue[0]
				activeOut = outQueue[0]
			}

			select {
			case request := <-s.in:
				requestQueue = append(requestQueue, request)
			case out := <-s.outChan:
				outQueue = append(outQueue, out)
			case activeOut <- activeRequest:
				requestQueue = requestQueue[1:]
				outQueue = outQueue[1:]
			}
		}
	}()
}
