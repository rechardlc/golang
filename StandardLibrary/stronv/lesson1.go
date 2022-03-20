package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// stronv: 实现了几本数据类和其字符串的互相转换

// strconv.Iton: 将int类型转化为string类型

func main() {
	var status = 200
	var toString = strconv.Itoa(status)
	fmt.Println(reflect.TypeOf(toString)) //  string
}
