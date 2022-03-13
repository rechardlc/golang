package main

import (
	"fmt"
	"reflect"
)

func main() {
	//a := '中'
	//b := "中"
	//var c rune = 'f'
	//fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b), c)
	//traversalString()
	changeString()
}

func traversalString() {
	var s string = "pprof.cn 富士康"
	for i := 0; i < len(s); i++ { // byte
		fmt.Println(s[i], reflect.TypeOf(s[i]))
	}
	fmt.Println("=========================")
	for _, i := range s { // rune
		fmt.Println(i, reflect.TypeOf(i))
	}
	fmt.Println(reflect.TypeOf(s))
}

/*
	对于字符串的修改，要先修改成[]rune/[]byte，发生修改，然后才能再次转化为string
*/
func changeString() {
	s1 := "hello"
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))
	s2 := "博客"
	runeS2 := []rune(s2)
	runeS2[0] = '服'
	fmt.Println(string(runeS2))

}
