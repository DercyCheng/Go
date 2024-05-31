package main

import (
	"fmt"
	"net"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {

		fmt.Println("error:", err)
		return
	}
	defer socket.Close()
	sendData := []byte("hello world")
	_, err = socket.Write(sendData)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)
}
