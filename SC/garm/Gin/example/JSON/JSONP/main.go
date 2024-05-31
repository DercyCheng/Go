package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/jsonp", func(c *gin.Context) {
		c.JSONP(http.StatusOK, gin.H{"wechat": "x"})
	})
	r.Run()
}
