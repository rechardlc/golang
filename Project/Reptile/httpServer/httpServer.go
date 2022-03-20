package httpServer

import (
	"example.com/m/v2/douBan"
	"example.com/m/v2/utils"
	"fmt"
	"log"
	"net/http"
)

func CreateServe() {
	http.HandleFunc("/", index)
	http.HandleFunc("/douban", douBan.DoubanMovie)
	ip, ok := utils.GetIPv4Addr().(string)
	if ok {
		fmt.Printf("start success:%v\n", ip+":8899")
		log.Fatal(http.ListenAndServe(ip+":8899", nil))
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello world!"))
}
