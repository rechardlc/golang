package main

import "fmt"

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
}
