package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	g1 := r.Group("/v1")
	{
		v1AdminGroup := g1.Group("/admin")
		{
			v1AdminGroup.GET("/users", func(c *gin.Context) {
				c.String(200, "/v1/admin/users")
			})
			v1AdminGroup.GET("/manager", func(c *gin.Context) {
				c.String(200, "/v1/admin/manager")
			})
			v1AdminGroup.GET("/photo", func(c *gin.Context) {
				c.String(200, "/v1/admin/photo")
			})
		}
		g1.GET("/test1", func(c *gin.Context) {
			c.String(http.StatusOK, "/v1/test1")
		})
		g1.GET("test2", func(c *gin.Context) {
			c.String(http.StatusOK, "/v1/test2")
		})
	}

	g2 := r.Group("/v2", func(c *gin.Context) {
		c.String(http.StatusOK, "v2中间件")
	})
	{
		g2.GET("/test1", func(c *gin.Context) {
			c.String(http.StatusOK, "/v2/test1")
		})
		g2.GET("/test2", func(c *gin.Context) {
			c.String(http.StatusOK, "/v2/test2")
		})
	}
	r.Run()
}
