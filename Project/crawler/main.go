package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
)

func main() {
	// 入口出解析第一个页面，珍爱网站
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.CityList,
	})
}
