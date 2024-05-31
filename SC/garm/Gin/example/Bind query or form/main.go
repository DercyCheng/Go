package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2021-10-26" time_utc:"1"`
}

func main() {
	r := gin.Default()
	r.GET("/start", startPage)
	r.Run()
}
func startPage(c *gin.Context) {
	// 如果是 `GET` 请求，只使用 `Form` 绑定引擎（`query`）
	// 如果是 `POST` 请求，首先检查 `content-type` 是否为 `JSON` 或 `XML`，然后再使用 `Form`（`form-data`）
	// https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48

	var person Person
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name, person.Address, person.Birthday)
	}
	c.String(200, "success")
}
