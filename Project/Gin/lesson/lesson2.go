package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"path/filepath"
	"time"
)

/*  关于HTML模板渲染测试
func main() {
	router := gin.Default()
	// filepath.Abs(path string)(string, error): 返回path所代表的绝对路径
	p, _ := filepath.Abs("templates")
	// router.LoadHTMLGlob(pattern string): 接受一个符合glob规则的路径
	router.LoadHTMLGlob(p + "/*")
	// router.GET(relativePath string, handlers ...HandlerFunc)
	router.GET("/", func(context *gin.Context) {
		// context.HTML(code int, name string, obj interface{})
		context.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Hello world",
		})
	})
	router.Run(":6688")
}
*/

func main() {
	r := gin.Default()
	// r.Delims(left, right sting): 自定义在模板中的需要解析的模板字符串
	r.Delims("{[{", "}]}")
	// r.SetFuncMap(funcMap template.FuncMap): 定义funcMap
	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	p, _ := filepath.Abs("templates/raw.html")
	r.LoadHTMLFiles(p)
	t := time.Now()
	y, m, d, h, mm, ss, ms := t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond()
	r.GET("/raw", func(c *gin.Context) {
		// 一大堆的验证逻辑
		c.HTML(http.StatusOK, "raw.html", gin.H{
			"now": time.Date(y, m, d, h, mm, ss, ms, time.UTC),
		})
	})
	if err := r.Run(":6688"); err != nil {
		return
	}
}
func formatAsDate(t time.Time) string {
	y, m, d := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", y, m, d)
}
