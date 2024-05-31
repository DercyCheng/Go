package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 路由重定向
	r.GET("/route", func(c *gin.Context) {
		c.Request.URL.Path = "/http"
		r.HandleContext(c)
	})

	// HTTP重定向
	r.GET("/http", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.bing.com")
	})

	r.Run()
}
