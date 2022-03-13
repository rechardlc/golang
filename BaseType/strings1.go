package main

import "fmt"

/*
	1. 字符串底层本质是一个[]byte, [:]可以用来切割字符串~
*/
func main() {
	str := "abd"
	split := str[:2]
	fmt.Println(split)
}
