package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginForm struct {
	UserName string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// 验证表单
//func main() {
//	r := gin.Default()
//	r.GET("/login", func(context *gin.Context) {
//		var form LoginForm
//		if context.ShouldBind(&form) == nil {
//			if form.UserName == "user" && form.Password == "password" {
//				context.JSON(http.StatusOK, gin.H{
//					"status": 200,
//					"data":   "login success",
//				})
//			} else {
//				context.JSON(http.StatusOK, gin.H{
//					"status": 400,
//					"data":   "login failed",
//				})
//			}
//		}
//	})
//	r.Run(":6688")
//}

func main() {
	r := gin.New()
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"index": "success",
		})
	})
	r.POST("/post", func(context *gin.Context) {
		id := context.Query("id")
		page := context.DefaultQuery("page", "1")
		fmt.Println(id, page)
		context.AsciiJSON(http.StatusOK, gin.H{
			"status": 200,
			"data":   "access Post method visit success!",
		})
	})
	r.GET("/get", func(context *gin.Context) {
		id := context.Query("id")
		page := context.DefaultQuery("page", "1")
		fmt.Println(id, page)
		context.AsciiJSON(http.StatusOK, gin.H{
			"status": 200,
			"data": map[string]interface{}{
				"id":   id,
				"page": page,
			},
		})
	})
	r.GET("/someJSON", func(context *gin.Context) {
		context.SecureJSON(http.StatusOK, []string{"dell", "richard"})
	})
	r.GET("/moreJSON", func(context *gin.Context) {
		context.JSON(http.StatusOK, struct {
			Name  string `json:"user"`
			Email string `json:"email"`
			Phone string `json:"phone"`
			Id    int    `json:"id"`
		}{
			"dell", "3300255192@qq.com", "13011111111", 10,
		})
	})
	r.Run(":6688")
}
