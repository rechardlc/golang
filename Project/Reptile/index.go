package main

import "example.com/m/v2/httpServer"

func main() {
	//douBan.Entry() // 爬取豆瓣网top250电影：同步方法
	//douBan.GoEntry() // 爬取豆瓣网top250电影: 异步方法
	httpServer.CreateServe()
}
