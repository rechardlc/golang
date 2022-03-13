package main

import (
	"fmt"
	"net"
)

func main() {
	// 建立一个udp链接，ip地址为0.0.0.0:2000
	udpConn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 20000,
	})
	if err != nil {
		return
	}
	defer func(udpConn *net.UDPConn) {
		if err := udpConn.Close(); err != nil {
			return
		}
	}(udpConn)
	sendData := []byte("Hello World!")
	// 通过Write方法发送数据
	if _, err = udpConn.Write(sendData); err != nil {
		fmt.Println("发送数据失败~")
		return
	}
	data := make([]byte, 4096)
	// 从连接中读取数据包，将数据包写入data中，返回读取的字节数和远程地址
	if n, remoteAddr, err := udpConn.ReadFromUDP(data); err != nil {
		fmt.Println("接受数据失败~", err)
	} else {
		fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)
	}
}
