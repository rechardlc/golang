package main

import (
	"fmt"
)

// panic: 表达很恐慌，完全不知道该干啥~
/*
	panic是一个很重的函数，要慎用
	执行：
	1. 会导致当前函数强制停止执行，当然也有一些例外，如http.HandleFunc中的异常就不挂掉程序，存在一定的保护机制
	2. 一直向上返回，执行每一层的defer
	3. 如果没有遇见recover，程序会退出
*/
// recover 只能在defer中使用
func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(err)
		}
	}()
	//panic(errors.New("this is an error "))
	//b := 0
	//a := 5 / b
	//fmt.Println(a)
	//panic("I don't know what to do")
}
func main() {
	tryRecover()
}
