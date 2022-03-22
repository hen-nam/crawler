package engine

import (
	"log"
)

// SimpleEngine 简单引擎
type SimpleEngine struct{}

// Run 运行
func (e SimpleEngine) Run(requests ...Request) {
	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]

		result, err := Worker(request)
		if err != nil {
			continue
		}

		requests = append(requests, result.Requests...)

		for _, item := range result.Items {
			log.Printf("Got item %+v", item)
		}
	}
}
