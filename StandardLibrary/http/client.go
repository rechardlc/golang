package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	resp, err := http.Get("https://www.imooc.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return
	}
	fmt.Printf(string(s))
}
