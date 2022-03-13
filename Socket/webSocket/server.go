package main

import (
	"chat/utils"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter() // 声明一个路由
	go utils.H.Run()          // 开启协程
	router.HandleFunc("/ws", utils.Myws)
	if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
		fmt.Println("err:", err)
	}
}
