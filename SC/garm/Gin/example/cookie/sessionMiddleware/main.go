package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	gsession "SC/garm/Gin/example/cookie/sessionMiddleware/gosession"
)

func main() {
	r := gin.Default()
	mgrObj, err := gsession.CreateSessionMgr(gsession.Redis, "localhost:6379")
	if err != nil {
		log.Fatalf("Create manager obj failed, err: %v\n", err)
		return
	}
	sm := gsession.SessionMiddleware(mgrObj, gsession.Options{
		Path:     "/",
		Domain:   "127.0.0.1",
		MaxAge:   120,
		Secure:   false,
		HttpOnly: true,
	})
	r.Use(sm)
	r.GET("/incr", func(c *gin.Context) {
		session := c.MustGet("session").(gsession.Session)
		fmt.Printf("%#v\n", session)
		var count int
		v, err := session.Get("count")
		if err != nil {
			log.Printf("get count from session failed, err: %v\n", err)
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.String(http.StatusOK, "count:%v", count)
	})
	r.Run()
}
