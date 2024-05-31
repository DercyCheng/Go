package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 自定义分隔符
	r.Delims("{[{", "}]}")

	// html渲染
	r.LoadHTMLGlob("E:/code/Go/src/SC/garm/Gin/example/templates/**/*")
	//r.LoadHTMLFiles("templates/")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})
	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})
	// 自定义的 html 模板渲染
	html := template.Must(template.ParseFiles("file1", "file2"))
	r.SetHTMLTemplate(html)

	r.Run()
}
