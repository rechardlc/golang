package main

import (
	"encoding/json"
	"fmt"
)

// 标签
/*
	结构体标签由一个或多个键值对组成。键与值使用冒号分隔，值用双引号括起来。
	键值对之间使用一个空格分隔。 注意事项： 为结构体编写Tag时，必须严格遵守键值对的规则。
	结构体标签的解析代码的容错能力很差，一旦格式写错，编译和运行时都不会提示任何错误，
	通过反射也无法正确取值。
	例如不要在key和value之间添加空格。
*/

type Student struct {
	ID           uint `json:"id"`
	Gender, name string
}

func main() {
	s1 := &Student{
		ID:     10,
		Gender: "女",
		name:   "dell",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json unmarshal failed!", err)
		return
	}
	fmt.Printf("%s\n", data)
	input()
	ce := make(map[int]Student)
	ce[1] = Student{1, "男", "dell"}
	ce[2] = Student{2, "女", "richard"}
	fmt.Println(ce)
	delete(ce, 2)
	fmt.Println(ce)
}

func change(ce []Student) {
	ce[0].name = "richard"
}
func input() {
	var ce []Student
	ce = []Student{
		Student{10, "女", "dell"},
		Student{20, "男", "bob"},
	}
	fmt.Println(ce, "before")
	change(ce)
	fmt.Println(ce, "after")
}
