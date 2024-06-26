package main

import (
	"fmt"
	"net"
)

// socket_stick/client/session.go

func main() {
	conn, err := net.Dial(
		"tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	msg := `Hello, Hello. How are you?`
	conn.Write([]byte(msg))
}
