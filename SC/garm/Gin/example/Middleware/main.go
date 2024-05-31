package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Gin框架允许开发者在处理请求的过程中 加入用户自己的钩子（Hook）函数
// 钩子函数就叫中间件，中间件适合处理一些公共的业务逻辑
// 比如登录认证、权限校验、数据分页、记录日志、耗时统计等
func main() {
	// gin默认中间件 默认使用了Logger和Recovery中间件
	//r := gin.Default()

	// 新建一个没有任何默认中间件的路由
	r := gin.New()
	// 注册一个全局中间件
	r.Use(StartCost())

	r.GET("/test", func(c *gin.Context) {
		name := c.MustGet("name").(string) // 从上下文取值
		log.Println(name)
	})

	// 为某个路由单独注册
	// 给/test2路由单独注册中间件（可注册多个）
	r.GET("/test2", StartCost(), func(c *gin.Context) {
		name := c.MustGet("name").(string) // 从上下文取值
		log.Println(name)
	})

	Group := r.Group("/group", StartCost())
	{
		Group.GET("/1", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "/group/1"})
		})
	}
	r.Run()
}

// StartCost 是一个统计耗时请求耗时的中间件
func StartCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "Tom") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		c.Next()
		// 不调用该请求的剩余处理程序
		// c.Abort()
		// 计算耗时
		cost := time.Since(start)
		c.JSON(http.StatusOK, cost)
	}
}
