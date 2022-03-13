package main

import "fmt"

// 获取输入~共三个函数：fmt.Scan/fmt.Scanf/fmt.Scanln
func main() {
	var (
		name    string
		age     uint8
		married bool
	)
	_, err := fmt.Scanln(&name, &age, &married)
	if err != nil {
		return
	}
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}
