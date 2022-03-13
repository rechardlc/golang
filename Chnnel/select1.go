package main

import (
	"fmt"
	"time"
)

func test1(ch chan string) {
	//time.Sleep(time.Second * 2)
	ch <- "test1"
}
func test2(ch chan string) {
	//time.Sleep(time.Second * 2)
	ch <- "test2"
}

// 多个channel是随机选择一个执行的
func main() {
	//out1, out2 := make(chan string), make(chan string)
	//for {
	//	go test1(out1)
	//	go test2(out2)
	//	select {
	//	case s1 := <-out1:
	//		fmt.Println("s1", s1)
	//	case s2 := <-out2:
	//		fmt.Println("s2", s2)
	//		//default:
	//		//	fmt.Println("nonono!")
	//	}
	//}
	testChFull()
}

func testChFull() {
	o1 := make(chan string, 10)
	go write(o1)
	for s := range o1 {
		fmt.Println("res:", s)
		time.Sleep(time.Second)
	}
}

func write(ch chan string) {
	for {
		select {
		case ch <- "hello":
			fmt.Println("write hello")
		default:
			fmt.Println("channel full")
			return
		}
		time.Sleep(time.Millisecond * 500)
	}
}
