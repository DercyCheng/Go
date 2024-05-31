package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Handle 只注册少量方法 一次性注册全部则使用 Any()
func Handle(r *gin.Engine, httpMethods []string, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	var routes gin.IRoutes
	for _, httpMethod := range httpMethods {
		routes = r.Handle(httpMethod, relativePath, handlers...)
	}
	return routes
}

// 符合RESTFUL 其实无所谓...
func main() {

	// 默认使用了2个中间件Logger()  Recovery()
	r := gin.Default()

	// GET用来获取资源
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET",
		})
	})

	// POST用来新建资源
	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST",
		})
	})

	// PUT用来更新资源
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT",
		})
	})

	// DELETE用来删除资源
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE",
		})
	})

	// 虽然这种方式比较便利，但是并不太推荐  因为他破坏了Restful 规范中HTTP Method的约束
	Handle(r, []string{"GET", "POST"}, "/", func(c *gin.Context) {})
	err := r.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}
