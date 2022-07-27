package main

import (
	"fmt"
)

func main() {
	createWays()
}

/*
	创建map2种方式
		1. make eg: make(map[string]int)
		2. 直接申明： map[string]int{}
*/
func createWays() {
	var way1 = map[string]int{"key": 1}
	way1["value"] = 10
	var way2 = make(map[string]string)
	way2["dell"] = "richard"
	fmt.Println(way1["key"], way2)
	// 判断是否存在key
	value, ok := way2["dell"]
	if ok {
		fmt.Println(way2["dell"], value, ok, "way2的值")
	} else {
		fmt.Println("way2中不存在dell的key", value == "", ok)
	}
	way2["age"] = "age"
	way2["name"] = "boll"
	// 通过range遍历map，类似于js中Object.entries
	// 遍历出来的结果是无序的
	for key, val := range way2 {
		fmt.Println(key, val)
	}
	// 通过delete删除key
	_, exist := way2["age"]
	fmt.Println(way2["age"], "exist:", exist)
	delete(way2, "age")
	_, del := way2["age"]
	fmt.Println("way2的age值：", way2["age"], del)
	// 创建元素为map类型的切片
	var way3 = make([]map[string]int, 0) //
	// 创建map值为切片
	var way4 = make(map[string][]int)
	way3 = append(way3, map[string]int{"age": 18, "head": 2})
	way4 = map[string][]int{"age": {1, 3, 4}}
	for i, v := range way3 {
		fmt.Println(i, v, "创建元素为map类型的切片")
	}
	for k, v := range way4 {
		fmt.Println(k, v, "创建map值为切片")
	}
	// 对应map的key排序问题，可以先将key放入keys的切片中，
	// 将keys排序，再通过rang循环~依次循环出map的结果~
	// 即map的排序结果
	way5 := make(map[string]int, 0)
	fmt.Println(len(way5))
	way5["age"] = 18
	fmt.Println(way5, len(way5))
	/*
		map的key
		1. map使用哈希表，必须可以比较相等
		2. 除了slice、map、function的内建类型都可以作为key
		3. struct类型不包含上述类型字段，也可以作为key
	*/
}
