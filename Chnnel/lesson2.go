package main

import "fmt"

// 构造通道作为参数~在多个任务函数间传递
func counter(out chan<- int) { // chan<- 只能发送的通道，不能接受
	for i := 0; i < 5; i++ {
		out <- i
	}
	close(out)
}
func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}
func printer(in <-chan int) { // <- chan只能接受的通道，不能发送
	for i := range in {
		fmt.Println(i)
	}
}
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}
