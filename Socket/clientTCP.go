package main

import (
	// bufio包实现了有缓冲的I/O。它包装一个io.Reader或io.Writer接口对象，
	// 创建另一个也实现了该接口，且同时还提供了缓冲和一些文本I/O的帮助函数的对象
	"bufio"
	"fmt"
	// net包提供了可移植的网络I/O接口，包括TCP/IP、UDP、域名解析和Unix域socket
	// 大部分使用者只需要Dial、Listen和Accept函数提供的基本接口；
	// 以及相关的Conn和Listener接口
	"net"
	// os包提供了操作系统函数的不依赖平台的接口,设计为Unix风格的
	"os"
	"strings"
)

func main() {
	// Dial的函数签名：func.md Dial(network, address string)(conn, error)
	conn, err := net.Dial("tcp", "127.0.0.1:2000") // 创建一个地址为127.0.0.1:2000的tcp链接
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	// 延迟服务~在return执行前执行关闭tcp链接，防止阻塞
	defer func(conn net.Conn) {
		if err := conn.Close(); err != nil {
			return
		}
	}(conn)
	// NewReader的函数签名：func.md NewReader(rd io.Reader) *Reader
	// 创建一个具有默认大小的缓冲，从r读取中的*Reader
	inputReader := bufio.NewReader(os.Stdin)
	for {
		// 读取标准输入的结果~，inputReader.ReadString接受一个分割字节符
		input, _ := inputReader.ReadString('\n')
		// strings.Trim去掉cutset包含的UTF-8码值的字符串(input字符串前后都去掉)
		inputInfo := strings.Trim(input, "\r\n")
		// strings.ToUpper将所有字母都转化为大写版本的拷贝
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}
		// 从连接中读取数据
		_, err := conn.Write([]byte(inputInfo))
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:]) // 返回的n字符长度
		fmt.Println(n, "n~~", len(buf))
		if err != nil {
			fmt.Println("recv failed err: ", err)
			return
		}
		fmt.Println(string(buf[:n]))
		fmt.Println(len(string(buf[:n])))
	}
}
