package main

import "fmt"

func main() {
	verifyQuote()
	declareSlice()
}

// 验证切片是数组的引用
func verifyQuote() {
	names := [4]string{"John", "paul", "George", "Ringo"}
	ages := [...]int8{1, 2, 4} // 数组的一种形式
	var head []int
	fmt.Printf("ages: %p\n", ages)
	fmt.Printf("ages: %p, bool: %v", head, head == nil)
	a := names[0:2] // 创建切片的方式
	b := names[1:3]
	fmt.Println(a, b, names, len(a), len(b), len(names), cap(a))
	b[0] = "dell"
	fmt.Println(a, b, names) // 切片是数组的引用，切片发生改变，数组也会发生改变
	/*
		1.b切片发生改变，a也发生改变~同时names数组也发生改变
	*/
}

/*
	1. slice在golang表现为一个动态数组
	2. 动态数组存在扩容行为，根据一定规则自动扩大
*/
func declareSlice() {
	var num1 []int // 声明长度为0、容量也为0的切片
	isNil := num1 == nil
	fmt.Println("num1:", num1, isNil, len(num1), cap(num1))
	var num2 = make([]int, 0, 3) // 声明长度为0，容量为3的切片
	fmt.Println("num2:", num2, num2 == nil, len(num2), cap(num2))
	num3 := make([]int, 3)
	fmt.Println("num3:", num3, num3 == nil, len(num3), cap(num3))
	num4 := make([]int, 0)
	fmt.Printf("num4:%p\n", num4)
	// 向slice中添加元素
	num4 = append(num4, 1, 2, 3, 4)
	fmt.Printf("num4 append after: %p\n", num4)
	// 验证数组拷贝
	num5 := [2]int{1, 2}
	num6 := num5
	num5[0] = 10
	isEqual := num6 == num5
	fmt.Printf("num5 内存地址：%p, num6 内存地址: %p\n", num5, num6)
	fmt.Println(isEqual)
	num7 := []int{1, 2, 3} // []中不加数字，创建的变量就是切片，加数字就是数组
	fmt.Println("num7:", num7, len(num7), cap(num7))
	/*
		创建切片的方式
			1. []string{} // []中为空
			2. var b = b[:]
			3. make([]int8, 0)
	*/
}
