package main

import (
	"io"
	"log"
	"net/http"
)

const form = `<html><body>
			<form action="#" method="post" name="bar">
            <input type="text" name="in">
            <input type="submit" value="Submit"></form>
			</body></html>`

type HandleFunc func(w http.ResponseWriter, r *http.Request)

func simpleserver(w http.ResponseWriter, r *http.Request) { io.WriteString(w, "<h1>hello world</h1>") }
func formserver(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch r.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		io.WriteString(w, r.FormValue("in"))
	}
}
func main() {
	http.HandleFunc("/test1", logPanic(simpleserver))
	http.HandleFunc("/test2", logPanic(formserver))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
func logPanic(function HandleFunc) HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", r.RemoteAddr, x)
			}
		}()
		function(w, r)
	}
}
