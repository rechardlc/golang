package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}
type Scheduler interface {
	Submit(Request)
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(out, e.Scheduler)
	}
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	count := 0
	for {
		result := <-out // 等待out的接受值，取out的值（主线程发生阻塞）
		for _, item := range result.Items {
			count++
			log.Printf("count: # %d, Got item: %v", count, item)
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

// 单通道channel： chan<-表示一个只向chan中发送数据，<-chan将chan中是数据发送出去
// @createWorker    createWorker
// @description   创建worker
// @auth      richard
// @param     in        chan Request         创建一个只接受的chan(只读通道)
// @param     out        chan ParseResult    创建一个只发送的chan(只写通道)
// @return    返回参数名        参数类型         "解释"
func createWorker(out chan ParseResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
