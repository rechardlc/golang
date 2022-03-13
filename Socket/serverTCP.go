package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer func(conn net.Conn) {
		if err := conn.Close(); err != nil {
			return
		}
	}(conn)
	for {
		// 新创建一个缓冲池
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据:", recvStr)
		if _, err = conn.Write([]byte(recvStr)); err != nil {
			return
		} // 发送数据
	}
}
func main() {
	// 监听tcp地址为127.0.0.1:2000地址
	listen, err := net.Listen("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println("listen failed, err: ", err)
		return
	}
	for { // 通过死循环，让main函数不结束~执行goroutine
		// listen.Accept(): Accept等待并返回下一个连接到该接口的链接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		// 创建一个协程
		go process(conn)
	}
}
