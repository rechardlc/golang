package main

import (
	"fmt"
)

const (
	a = 1
	b = 10
	c = iota + 2
	_
	d
)

func main() {
	//a := 1 + 2i
	//fmt.Println(cmplx.Abs(a))
	fmt.Println(a, b, c, d)
}
