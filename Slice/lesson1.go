package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func main() {
	var s1 []int // 创建nil切片
	if s1 == nil {
		fmt.Println("is nil")
	} else {
		fmt.Println("is not nil")
	}
	var s2 = make([]int, 0)
	if s2 == nil {
		fmt.Println("is nil")
	} else {
		fmt.Println("is not nil")
	}
	var s3 = []int{1, 2}
	var s4 = s3[0:1:2]
	fmt.Println(s4)
	var sliceToString = []int{1, 2, 3, 4}
	var successToString = strings.Replace(strings.Trim(fmt.Sprint(sliceToString), "[]"), "", "", -1)
	fmt.Println(reflect.TypeOf(successToString), successToString)
	// 从slice中获取一块内存指针
	s5 := make([]byte, 200)
	prt := unsafe.Pointer(&s5[0])
	fmt.Println(prt)

	// 验证切片中每一元素的地址
	fmt.Println("============分割=============")
	s6 := []map[string]int{
		{"age": 12},
		{"age": 19},
		{"age": 20},
	}
	// range会发生复制现象，导致出现值传递现象，值传递拷贝副本~始终指向一个地址
	for _, v := range s6 {
		fmt.Printf("%p\n", &v)
	}
}
