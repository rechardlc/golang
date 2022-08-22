package main

import (
	"fmt"
)

type Test struct {
	name string
}

// Close defer也会导致值拷贝
func (t Test) Close() {
	fmt.Println(t.name, "closed")
}

func main() {
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		defer t.Close() // a / b / c
	}
	n1, n2 := 10, 20
	defer func() {
		fmt.Println(n1, "n1-defer")
	}()
	defer func() {
		fmt.Println(n2, "n2-defer")
	}()
	n1 = n2 + n1
	fmt.Println(n1, "done")
	inputParams("a", "b", "c", "d", "f", "g")
}

type N interface{}

// 多参数传递
func inputParams(n ...string) {
	input := [...]string{"h", "i", "j", "k", "l", "m"}
	j := 0
	for len(n) > 0 {
		i := n[0]
		n = n[1:]
		fmt.Printf("队列头=%s， 剩余队列数据=%s\n", i, n)
		if j < len(input) {
			n = append(n, input[j])
			j++
		}
	}
}
