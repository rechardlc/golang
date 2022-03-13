package main

import (
	"fmt"
	"net/http"
)

func main() {
	//http.HandleFunc: 注册一个处理函数handler和对应的模式
	http.HandleFunc("/go", myHandler)
	// 监听tcp地址~
	err := http.ListenAndServe("127.0.0.1:8000", nil)
	if err != nil {
		return
	}
}

// http.ResponseWriter: ResponseWriter接口被http处理器用于构造http回复
// http.Request: Request类型代表一个服务器接收到的或者客户端发送出去的http请求
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "链接成功")
	fmt.Println("method:", r.Method)
	fmt.Println("url:", r.URL.Path)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)
	// http.ResponseWriter.Write: 向链接中写入数据
	if write, err := w.Write([]byte("www.51mh.com")); err != nil {
		return
	} else {
		fmt.Println(write)
	}
}
