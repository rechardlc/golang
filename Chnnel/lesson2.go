package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 构造通道作为参数~在多个任务函数间传递
func counter(out chan<- int, count int) { // chan<- 只能发送的通道，不能接受, 只写通道
	for i := 0; i <= count; i++ {
		out <- i // 将i发送给out
	}
	close(out)
}
func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}
func printer(in <-chan int) { // <- chan只能接受的通道，不能发送，只读通道
	for i := range in {
		fmt.Printf("from in<-chan: %d\n", i)
	}
}

func getRandChan() <-chan int {
	ch1 := make(chan int)
	go func() {
		rand.Seed(time.Now().Unix())
		for {
			time.Sleep(time.Millisecond * 10)
			ch1 <- rand.Intn(10) // 返回一个[0,10]的随机数
		}
	}()
	return ch1
}

func main() {
	out := make(chan int) // out
	in := make(chan int)  // in
	go counter(out, 10)
	go squarer(in, out)
	printer(in)

	//ch := getRandChan()
	//for {
	//	af := <-ch
	//	fmt.Printf("%d\n", af)
	//}
}
