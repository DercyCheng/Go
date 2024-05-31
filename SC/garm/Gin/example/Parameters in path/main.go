package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Gin的路由是单一的 不能有重复 注册了/users/:id  不能再注册匹配/users/:id模式的路由
func main() {
	r := gin.Default()

	// : 匹配的是  :  后的字段
	// * 匹配的是  所有字段
	r.GET("/users/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
	r.POST("/users/:name/*action", func(c *gin.Context) {
		b := c.FullPath() == "/users/:name/*action" // true
		c.String(http.StatusOK, "%t", b)
	})

	r.GET("/users/group", func(c *gin.Context) {
		c.String(http.StatusOK, "the available group is [...]")
	})

	//获取Get参数  匹配格式 : /welcome?firstname=Jane&lastname=Doe
	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "first")
		lastname := c.Query("lastname")
		//是 c.Request.URL.Query().Get("lastname") 的简写
		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	//获取Post参数 postman使用post
	r.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		// 如果get不到值就返回默认值
		nick := c.DefaultPostForm("nick", "anonymous")
		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	// GET和POST结合 url: /post?id=1234&page=1 HTTP/1.1
	r.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})

	//获取form参数
	r.POST("/user/search", func(c *gin.Context) {
		// DefaultPostForm取不到值时会返回指定的默认值
		//username := c.DefaultPostForm("username", "小王子")
		username := c.PostForm("username")
		address := c.PostForm("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	// 获取json参数
	r.POST("/json", func(c *gin.Context) {
		b, _ := c.GetRawData()
		var m map[string]interface{}
		json.Unmarshal(b, &m)
		c.JSON(http.StatusOK, m)
	})

	// queryArray()
	// url:http://localhost:8080/array/?media=a&media=b
	r.GET("/array", func(c *gin.Context) {
		c.JSON(200, c.QueryArray("media"))
	})

	// queryMap()
	// url:http://localhost:8080/map?ids[a]=123&ids[b]=456&ids[c]=789
	r.GET("/map", func(c *gin.Context) {
		c.JSON(http.StatusOK, c.QueryMap("ids"))
	})
	r.Run()
}
