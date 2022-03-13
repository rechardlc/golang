package main

import (
	"fmt"
)

// FormatFunc 定义函数类型: 函数签名
type FormatFunc func(s string, x, y int) string
type Person struct {
	Name string
	Age  uint8
	do   FormatFunc
}

func main() {
	var p1 = &Person{
		Name: "richard",
		Age:  19,
		do: func(s string, x, y int) string {
			fmt.Println(rune(x), "string(rune(x))")
			return s + "" + string(rune(x)) + string(rune(y))
		},
	}
	result := p1.do("move", 1, 2)
	fmt.Println(result)
	var f2 = (func(x int) int { return 10 * x })(10)
	fmt.Println(f2)
	c := Closure()
	c()
	c()
	c()
	c()
	TestDefer()
}

func Closure() func() int {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}

func TestDefer() {
	fmt.Println("======split=======")
	whatever := [5]int{}
	for i := range whatever {
		defer fmt.Println(i) // 4/3/2/1/0 栈执行顺序(先进后出)
	}
}
