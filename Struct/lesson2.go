package main

import (
	"fmt"
)

/*
	声明结构体，大写开头字母，可以公开访问，小写开头字母只能在当前包中访问
*/
// Person 定义Person结构体
type Person struct {
	name, city string
	age        uint8
}

func main() {
	p1 := constructorFunc("dell", "成都", 12)
	fmt.Println(p1, p1.city, p1.name, p1.age)
	p1.Dream()
	fmt.Println(p1.age)
	p1.SetAge(100)
	fmt.Println(p1.age)
	// 继承测试
	var d1 = &Dog{
		feet: 10,
		Animal: &Animal{
			AName: "来福",
		},
		Person: &Person{
			name: "dell",
			age:  19,
			city: "成都",
		},
	}
	d1.move()
	d1.barking()
	d1.Dream()
	d1.SetAge(100)
}

// Go语言没有构造函数，需要自己通过struct实现
func constructorFunc(name, city string, age uint8) *Person {
	return &Person{
		name: name,
		city: city,
		age:  age,
	}
}

// Dream 类似于js中function prototype
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言\n", p.name)
}
func (p *Person) SetAge(age uint8) {
	p.age = age
}

// 通过struct实现其他变成语言中继承

type Animal struct {
	AName string
}

func (a *Animal) move() {
	fmt.Printf("名为%v的动物会动~\n", a.AName)
}

type Dog struct {
	feet    int8
	*Animal // 通过嵌套匿名结构体
	*Person
}

func (d Dog) barking() {
	fmt.Printf("名为%v的狗狗叫了%v声\n", d.AName, d.feet)
}
