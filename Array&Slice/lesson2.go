package main

import "fmt"

/*
	在Golang中，如果将一个数组传入某个函数，
	在函数内部接收到的是数组的副本而非指针。如果想传入指针，
	有两种做法:
	一是显式加上指针
	第二就是传递切片（推荐做法）
*/
func main() {
	elemAddTest()
	elemAddByPointTest()
	elemAddBySliceTest()
	interview()
	sliceTest()
}

func elemAdd(arr [3]int) {
	for idx := range arr {
		arr[idx] += 1
	}
	fmt.Println("new array:", arr) // [2, 3, 4]
}

func elemAddTest() {
	num1 := [3]int{1, 2, 3}
	elemAdd(num1)
	fmt.Println(num1) // [1, 2, 3]
}

// 1.显式加上指针
func elemAddByPoint(arr *[3]int) {
	for idx, _ := range arr {
		arr[idx] += 1
	}
	fmt.Println("new", arr) // &[2, 3, 4]
}

func elemAddByPointTest() {
	num1 := &[3]int{1, 2, 3}
	elemAddByPoint(num1)
	fmt.Println(num1) // [2, 3, 4]
}

// 2.传递切片
func elemAddBySlice(arr []int) {
	for idx, _ := range arr {
		arr[idx] += 1
	}
	fmt.Println("new slice:", arr) // [2, 3, 4]
}

func elemAddBySliceTest() {
	num1 := []int{1, 2, 3}
	elemAddBySlice(num1)
	fmt.Println(num1) // [2, 3, 4]
}

func interview() {
	s := make([]int, 5) // 切片~
	s = append(s, 1, 2, 3)
	// 第一次打印
	fmt.Println(s) // [0, 0, 0, 0, 0, 1, 2, 3]

	var fb = func(arr []int) {
		arr[0] = 10
		arr = append(arr, 4)
	}
	fb(s)
	// 第二次打印
	fmt.Println(s) // [10,0,0,0,0,1,2,3,4]
}
func sliceTest() {
	// 声明两个切片，一开始两个切片的地址是不一样的
	var t = make([]int, 0, 10)
	var s = make([]int, 0, 10)

	fmt.Printf("addr:%p len:%v content:%v\n", t, len(t), t) // addr:0xc00009a000       len:0 content:[]
	fmt.Printf("addr:%p len:%v content:%v\n", s, len(s), s) // addr:0xc00009a050       len:0 content:[]

	t = append(s, 1, 2, 3, 4)
	fmt.Printf("addr:%p len:%v content:%v\n", t, len(t), t) // addr:0xc00009a050       len:4 content:[1 2 3 4]
	fmt.Printf("addr:%p len:%v content:%v\n", s, len(s), s) // addr:0xc00009a050       len:0 content:[]

	var m = make([]int, 0, 10)
	var n = append(m, 1, 2)
	fmt.Println(m, n)

}
