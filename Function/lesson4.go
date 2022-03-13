package main

import (
	"fmt"
)

func main() {
	// 1、1、2、3、5、8、13、21、34、55、89
	r := fbnq(11, []int{1, 1}, 2)
	fmt.Println(r)
	rFor, i := fibona(), 1
	for {
		i++
		if i > 10 {
			fmt.Println(rFor())
			break
		}
		rFor()
	}
	fmt.Println(f(3))
	fmt.Println(peach(2))
}
func fbnq(n int, x []int, i int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	if n == i {
		return x[1]
	}
	x[0], x[1] = x[1], x[0]+x[1]
	i++
	return fbnq(n, x, i)
}

func fibona() func() int {
	i, j := 1, 1
	return func() int {
		i, j = j, j+i
		return j
	}
}

func f(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*f(n-1) + 1
		// n = 4
		// 2*f(2)+1 // 2*(2*(2*(1) + 1) + 1) + 1
	}
}

// 应该算第N有多少桃子？peach(n) = (peach(n + 1) + 1) * 2
// 1: x - (x/2 + 1) // peach(n) = (peach(n+1) + 1) * 2
func peach(n int) int {
	if n == 10 {
		return 1
	}
	return (peach(n+1) + 1) * 2
}
