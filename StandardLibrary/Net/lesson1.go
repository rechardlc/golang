package main

import (
	"net"
)

func main() {

}

// GetLocalIpAddr
//  @Description: 获取本地IP地址
//  @return interface{}
//
func GetLocalIpAddr() interface{} {
	addrs, _ := net.InterfaceAddrs()
	for _, address := range addrs {
		// address.(*net.IpNet)：判断address是否为net.IpNet类型
		// ipNet.Ip.IsLoopback()：判断ip是环回地址
		// 声明ipNet类型，并且不是Ipv4的环回地址(127.0.0.1)
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			// 将Ipv4转化为4字节表示
			if ip := ipNet.IP.To4(); ip != nil {
				return ipNet.IP.String()
			}
		}
	}
	return nil
}
