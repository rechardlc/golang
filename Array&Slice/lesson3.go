// 验证append、copy方法
package main

import (
	"fmt"
)

func main() {
	verifyCopy()
	verifyAppend()
	f := Filter([]int{1, 2, 3}, func(i int) bool {
		if i > 0 {
			return true
		}
		return false
	})
	fmt.Println(f)
	p := []byte{2, 3, 5}
	p = AppendByte(p, 7, 11, 13)
	fmt.Println(p)
}
func verifyCopy() {
	a := []int{1, 2, 3}
	copy(a, []int{11, 22, 33, 44, 55})
	fmt.Println(a) // [11,22,33] 因为a的容量为3，只能复制3个
}
func verifyAppend() {
	a := []int{1, 2, 3}
	a = append(a, []int{11, 22, 33}...) // append切片的时候，使用...语法
	fmt.Println(a)
}
func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)     // slice的长度
	n := m + len(data)  // slice长度 + data长度
	if n > cap(slice) { // 如果n大于slice的容量
		newSlice := make([]byte, (n+1)*2) // 创建一个新的切片，长度为2倍N+1
		copy(newSlice, slice)             // 将slice复制到newSlice中
		slice = newSlice                  // slice还原newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
func Filter(s []int, fn func(int) bool) []int {
	var p []int
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}
