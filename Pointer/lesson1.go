package main

import (
	"fmt"
)

func main() {
	var p *string
	fmt.Println(p)
	fmt.Printf("p的值是%v\n", p)
	if p == nil {
		fmt.Printf("p is nil: %p\n", p)
	} else {
		fmt.Printf("p is not nil: %v\n", p)
	}
	TestNil()
}

// TestNil 创建指针必须要有创建对应的内存~否者将会出现panic
func TestNil() {
	var a *int
	a = new(int)
	*a = 10
	fmt.Println(*a)
}
