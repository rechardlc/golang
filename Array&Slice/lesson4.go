package main

import "fmt"

func main() {
	//verifyArrayIsValueType()
	verifyAppend()
}
func verifyArrayIsValueType() {
	// 由于数组是值类型，所以可以当做map的key
	var s1 = make([]int, 16, 32)
	fmt.Println(s1)
	var m1 = make(map[interface{}]interface{})
	var a1 = [2]int{1, 2}
	m1[a1] = a1
	m1[s1] = s1
	fmt.Println(m1)
}
func verifyAppend() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6] // 2,3,4,5 len:4, cap: 6
	s2 := s1[3:5]  // 3,4 len: 2, cap: 3
	fmt.Println(s1, len(s1), cap(s1), arr)
	fmt.Println(s2, len(s2), cap(s2), arr)
	s3 := append(s2, 10)
	fmt.Println(s3, len(s3), cap(s3), arr)
	s4 := append(s3, 11)
	fmt.Println(s4, len(s4), cap(s4), arr) // [5 6 10] 3 3 [0 1 2 3 4 5 6 10]
	s5 := append(s4, 12)                   // [5 6 10 11] 4 6 [0 1 2 3 4 5 6 10]
	fmt.Println(s5, len(s5), cap(s5), arr) // [5 6 10 11 12] 5 6 [0 1 2 3 4 5 6 10]
	// 总结：1. slice只能向后
}
