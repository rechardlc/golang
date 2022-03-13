package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://127.0.0.1:8000/go")
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			return
		}
	}(resp.Body)
	fmt.Println(resp.Status)
	fmt.Println(resp.Header)
	for {
		var buf = make([]byte, 1024)
		if n, err := resp.Body.Read(buf); err != nil && err != io.EOF {
			fmt.Println(err)
			return
		} else {
			fmt.Println("读取完毕")
			res := string(buf[:n])
			fmt.Println(res)
			break
		}
	}
}
