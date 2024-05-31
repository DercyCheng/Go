package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		// 美化json
		c.IndentedJSON(http.StatusOK, gin.H{"message": "hello world"})
	})
	r.GET("/user", func(c *gin.Context) {
		var user struct {
			id   string `json:"id"`
			name string `json:"name"`
			age  int    `json:"age"`
		}
		user.id = "001"
		user.name = "tom"
		user.age = 18
		c.PureJSON(http.StatusOK, user)
	})
	// 使用 JSON 生成具有转义的非 ASCII 字符的 ASCII-only JSON
	// gin.H可以直接代替
	r.GET("/asciijson", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "go",
			"tag":  "<br>",
		}
		// {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})

	r.Run()
}
