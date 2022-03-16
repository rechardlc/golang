package main

import (
	"fmt"
	"regexp"
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

// 闭包小测试: 接受一个文件后缀名(比如:.jpg),若存在后缀，返回文件名，若不存在后缀，加上后缀返回~
type suffixFunc func(fileName string) string

func makeSuffix(suffix string) suffixFunc {
	//re, _ := regexp.Compile(suffix + "$")
	return func(fileName string) string {
		// 实现方式：1、正则 2.strings.HasSuffix函数
		//if ok := strings.HasSuffix(fileName, suffixInner); ok {
		//	return fileName
		//}
		//return fileName + suffixInner
		//if ok := re.MatchString(fileName); ok {
		//	return fileName
		//}
		//return fileName + suffix

		if ok, _ := regexp.MatchString(suffix, fileName); ok {
			return fileName
		}
		return fileName + suffix
	}
}
