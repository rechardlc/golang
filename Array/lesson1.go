package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 数组声明方式
var arr1 = [5]int{1, 2, 3}
var arr2 = [...]int{1, 2, 3}
var arr3 = [...]string{3: "hello", 4: "tom"}
var arr4 = [...]struct {
	name string
	age  uint8
}{
	{"user1", 10},
	{"user2", 20},
}

var arr5 = [...][2]int{{1, 2}, {23, 00}}
var arr6 = [2]int{1, 2}
var arr7 = [2]int{1, 2}

func main() {
	fmt.Println("arr6 == arr7:", arr6 == arr7) // true :  1、数组长度也是类型 2、数组的长度是值类型可以直接==
	fmt.Println(arr3, len(arr3), arr2, len(arr2), arr1, len(arr1), arr4, arr5)
	var arr6 [5]int
	printArr(&arr6)
	fmt.Println(arr6)
	var arr7 = [...]int{1, 2, 3, 4, 5}
	printArr(&arr7)
	fmt.Println(arr7)
	var b [10]int
	rand.Seed(time.Now().Unix())
	for i, _ := range b {
		b[i] = rand.Intn(1000)
	}
	sum := sumArr(b)
	fmt.Println(sum)
}

func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

// 数组求和
func sumArr(arr [10]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	fmt.Println(arr)
	return sum
}
