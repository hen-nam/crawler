package engine

import (
	"log"
)

// ConcurrentEngine 并发引擎
type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	Processor   Processor
	ItemChan    chan Item
}

// Scheduler 调度器
type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

// ReadyNotifier 准备就绪通知器
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

// Processor 请求处理器
type Processor func(Request) (ParseResult, error)

// Run 运行
func (e *ConcurrentEngine) Run(requests ...Request) {
	out := make(chan ParseResult)

	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		in := e.Scheduler.WorkerChan()
		e.makeWorker(in, out)
	}

	for _, request := range requests {
		if isDuplicate(request.Url) {
			log.Printf("Duplicate url: %s", request.Url)
			continue
		}
		e.Scheduler.Submit(request)
	}

	for {
		result := <-out
		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				log.Printf("Duplicate url: %s", request.Url)
				continue
			}
			e.Scheduler.Submit(request)
		}
		for _, item := range result.Items {
			go func(i Item) {
				e.ItemChan <- i
			}(item)
		}
	}
}

// makeWorker 创建工作器
func (e *ConcurrentEngine) makeWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			e.Scheduler.WorkerReady(in)
			request := <-in
			result, err := e.Processor(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

// 已访问的地址
var visitedUrls = make(map[string]bool)

// isDuplicate 检查地址是否重复出现
func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
