package main

import (
	"github.com/gin-gonic/gin"
)

type A struct {
	FieldA string `form:"field_A"`
}
type B struct {
	//嵌套结构体
	NestedStruct A
	FieldB       string `form:"field_B"`
}
type C struct {
	//结构体指针
	NestedStuctPointer *A
	FieldC             string `form:"field_C"`
}
type D struct {
	//匿名结构体
	NestedAnonyStruct struct {
		FieldX string `form:"field_X"`
	}
	FieldD string `form:"field_D"`
}

func getDB(c *gin.Context) {
	var b B
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}
func getDC(c *gin.Context) {
	var b C
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStuctPointer,
		"c": b.FieldC,
	})
}

func getDD(c *gin.Context) {
	var b D
	c.Bind(&b)
	c.JSON(200, gin.H{
		"x": b.NestedAnonyStruct,
		"c": b.FieldD,
	})
}

//localhost:8080/getb?field_A=hello&field_B=world&field_C=?&field_D=!
func main() {
	r := gin.Default()
	r.GET("/getb", getDB)
	r.GET("/getc", getDC)
	r.GET("/getd", getDD)
	r.Run()
}
