package main

import (
	"fmt"
	"reflect"
)

// MyInt 通过type字段来定义新的类型
// 判断是否为自定义类型还是类型别名：
// 看是否存在等号~存在为类型别名，不存在为自定义类型
type MyInt int      // 声明一个MyInt类型，该类型具有int属性
type AliasInt = int // 声明AliasInt为int属性的别名

// Person 定义Person类型
type Person struct {
	name, city string
	age        uint8
}

// 通过war方法定义匿名结构体
var user struct {
	name string
	age  uint8
}

func main() {
	fmt.Printf("MyInt=%+v\n", MyInt(12))
	fmt.Println(reflect.TypeOf(MyInt(12)), "MyInt 类型")
	fmt.Println(reflect.TypeOf(AliasInt(12)), "Alies -> Int")
	var p1 = Person{age: 18, name: "richard"}
	fmt.Println(p1.age, p1.city == "", p1.name)
	user.age = 12
	user.name = "dell"
	fmt.Println(user)
	var p2 = new(Person) // 通过new函数生成一个内存地址
	fmt.Printf("p2->指针：%T, 内存地址：%p, 完整结构:%#v\n", p2, p2, p2)
	p2.city = "华阳"
	fmt.Println(p2.city)
	p3 := &Person{} // 通过&(指针取值)操作，相当于new行为
	p3.age = 12     // 底层实现了(*p3).age = 12，Go实现p3 -> (*p3)的语法糖
	fmt.Printf("p3->指针：%T, 内存地址：%p, 完整结构:%#v\n", p3, p3, p3)

	interview1()
}

func interview1() {
	// 声明一个Student的结构体
	type Student struct {
		name string
		age  uint8
	}
	m := make(map[string]*Student) // m为Student结构的指针(内存地址)
	fmt.Printf("stus:%T==%+v==%p\n", m, m, m)
	stus := []Student{ // 实例化一个stus的切片
		{name: "richard", age: 10},
		{name: "dell", age: 24},
		{name: "bob", age: 100},
	}
	fmt.Printf("stus:%T==%+v==%p\n", stus, stus, stus)
	//for _, stu := range stus { // 循环stus切片
	for i := 0; i < len(stus); i++ { // 循环stus切片
		m[stus[i].name] = &stus[i] // m["richard"] = &{name: richard, age: 10}
		fmt.Printf("stus:%T==%+v==%p\n", &stus[i], &stus[i], &stus[i])
		fmt.Println(m[stus[i].name])
		/*
				使用range导致的现象
			    stus:*main.Student==&{name:richard age:10}==0xc0000a40c0
				&{richard 10}
				stus:*main.Student==&{name:dell age:24}==0xc0000a40c0
				&{dell 24}
				stus:*main.Student==&{name:bob age:100}==0xc0000a40c0
				&{bob 100}
		*/
	}
	// stus:map[string]*main.Student==map[bob:0xc0000a40c0 dell:0xc0000a40c0 richard:0xc0000a40c0]==0xc000098240
	// 为什么会指向相同地址？// range会复制值，涉及到值传递现象
	fmt.Printf("stus:%T==%+v==%p\n", m, m, m)

	for k, v := range m {
		fmt.Println(k, "=>", v.name) // name,
	}
}
