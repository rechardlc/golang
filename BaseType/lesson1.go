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
	//indexToByte()
	//var temp = "fdsfsfds"
	//var t = temp[0:3]
	//fmt.Println(temp, t)
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
	s1 := "hello world"
	fmt.Println([]rune(s1)) // rune uint64
	byteS1 := []byte(s1)    // byte uint8
	fmt.Println(byteS1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))
	s2 := "博客"
	runeS2 := []rune(s2)
	runeS2[0] = '服'
	fmt.Println(string(runeS2))

}

//
//  indexToByte
//  @Description: 通过索引取字符串的值，取出来的值类型为uint8(byte)类型,而非string
//
func indexToByte() {
	var s = "hello word"
	fmt.Println(reflect.TypeOf(s))    // string
	fmt.Println(reflect.TypeOf(s[5])) // uint8(byte)
}
