package main

import (
	"fmt"
	"net"
)

//不需要建立连接就能直接进行数据发送和接收
//属于不可靠的、没有时序的通信 但是UDP协议的实时性比较好
//通常用于视频直播相关领域

// net包实现
func main() {
	//开辟UDP
	socket, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer socket.Close()
	//不需要建立连接,直接发送数据
	for {
		//创建空间存放数据
		var data [1024]byte
		n, addr, err := socket.ReadFromUDP(data[:]) //接收数据
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
		fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
		_, err = socket.WriteToUDP(data[:], addr)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
	}

}
