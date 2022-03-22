package scheduler

import "crawler/crawler/engine"

// SimpleScheduler 简单调度器
type SimpleScheduler struct {
	out chan engine.Request
}

// Submit 提交请求
func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() {
		s.out <- request
	}()
}

// WorkerChan 获取连接工作器的通道
func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.out
}

// WorkerReady 准备连接工作器的通道
func (s *SimpleScheduler) WorkerReady(out chan engine.Request) {}

// Run 运行
func (s *SimpleScheduler) Run() {
	s.out = make(chan engine.Request)
}
