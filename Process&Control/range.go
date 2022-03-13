package main

import "fmt"

// range 会复制对象
func main() {
	copyObject()
}
func copyObject() {
	a := [3]int{1, 2, 3}
	for i, v := range a { // i, v 都是从复制品中取出的
		if i == 0 {
			a[1], a[2] = 99, 999
			fmt.Println(a)
		}
		a[i] = v + 100
	}
	fmt.Println(a)
}
