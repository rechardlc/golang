package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type intGen func() int // 声明一个intGen类型

func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// 在go语言除了内置的简单类型外，都可以定义方法
// 在intGen类型实现Read接口
// todo 对于接口实现，并且调用逻辑~还没有理解清楚，需要再次理解
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci()
	printFileContents(f)
}
