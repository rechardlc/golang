package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 基础api测试
	// gin.Default()：生成一个默认gin引擎
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO",
			"tag":  "<br>",
		}
		// context.AsciiJSON(code int, obj interface{})：将生成具有转义的非 ASCII 字符的 ASCII-only JSON
		context.AsciiJSON(http.StatusOK, data)
	})
	if err := r.Run(); err != nil {
		fmt.Println("")
		return
	}
}
