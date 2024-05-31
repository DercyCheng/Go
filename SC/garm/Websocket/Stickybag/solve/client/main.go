package main

import (
	"SC/garm/Websocket/Stickybag/solve/proto"
	"fmt"
	"net"
)

// https://www.liwenzhou.com/posts/Go/15_socket/#autoid-2-3-2
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 10; i++ {
		msg := `Hello, Hello. How are you?`
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(data)
	}
}
