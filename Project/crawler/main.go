package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main() {
	// 入口出解析第一个页面，珍爱网站
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}
	//e := engine.ConcurrentEngine{
	//	Scheduler:   &scheduler.SimpleScheduler{},
	//	WorkerCount: 100,
	//}
	//e := engine.SimpleEngine{}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.CityList,
	})
}
