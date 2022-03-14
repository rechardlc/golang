package utils

import (
	"fmt"
	"net"
)

func GetIPv4Addr() interface{} {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if addrs, err := net.InterfaceAddrs(); err != nil {
		return err
	} else {
		for _, addr := range addrs {
			if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
				if ip := ipNet.IP.To4(); ip != nil {
					return ipNet.IP.String()
				}
			}
		}
	}
	return nil
}
