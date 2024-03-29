/*
	go语言接口类似于duck typing形式，在go语言中称为结构化类型系统
	如何实现接口：一个对象只要全部实现了接口中的方法，那么就实现这个接口

	注意：为了保护Go语言的职业生涯，明确记住interface是一种类型
*/
package main

import (
	"fmt"
)

func main() {
	var x Sayer
	a := cat{}
	x = a
	x.say()
	TestDone()
	judgeType("dog")
}

type Sayer interface {
	say()
}
type dog struct{}
type cat struct{}

func (c cat) say() {
	fmt.Println("旺旺")
}

type People interface {
	Speak(string) string
	Eat() bool
}
type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是SB"
	} else {
		talk = "您好啊~"
	}
	return
}
func (stu Student) Eat() bool {
	return false
}
func TestDone() {
	var peo People = &Student{}
	think := "sb"
	fmt.Println(peo.Speak(think))
}

// 空接口：接受任意值：如~func.md(x Interface{})Interface{}/make(map[string]Interface{})/....
// x.(T),判断T是否是实现了x接口，所以x必须为接口类型
// 通过类型断言判断~空接口对应的值 x.(Type):
func judgeType(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("x is a string, value is %v\n", v)
	case bool:
		fmt.Printf("x is a bool, value is %v\n", v)
	default:
		fmt.Printf("x is a custom type, value is %v\n", v)
	}

}
