package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 20000,
	})
	if err != nil {
		fmt.Println("listen failed err:", err)
		return
	}
	if err := listen.Close(); err != nil {
		return
	}
	for {
		var data [1024]byte
		if n, addr, err := listen.ReadFromUDP(data[:]); err != nil {
			fmt.Println("read udp failed, err:", err)
			continue
		} else {
			fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
			_, err = listen.WriteToUDP(data[:n], addr)
			if err != nil {
				fmt.Println("write to upd failed, err:", err)
				continue
			}
		}
	}
}
