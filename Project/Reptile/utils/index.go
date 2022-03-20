package utils

import (
	"fmt"
	"net"
)

// GetIPv4Addr
//  GetIPv4Addr
//  @Description: 获取IP地址
//  @return interface{}
//
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

//
//  removeSliceRePeat
//  @Description: 数组去重
//  @param s
//  @return []interface{}
//
func removeSliceRePeat(s []interface{}) []interface{} {
	m := make(map[interface{}]interface{})
	rs := make([]interface{}, 0)
	for _, v := range s {
		if _, ok := m[v]; !ok {
			m[v] = v
			rs = append(rs, v)
		}
	}
	return rs
}
