package main

import "fmt"

// channel是一种引用类型
/*
	1. 通道建立~指定容量大小，就是有缓冲channel，否则就是无缓冲channel
	2. 无缓冲通道会造成阻塞，有被称为同步通道
*/
/*  channel 基础用法
var ch1 chan int
var ch2 chan bool
var ch3 chan []interface{}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err, "捕获panic")
		}
	}()
	ch4 := make(chan int) // 声明一个channel类型
	go receiver(ch4)      // 发起一个协程(async)
	//close(ch4)
	ch4 <- 10             // 将10发送给ch4通道(channel)中
	fmt.Println("发送成功！")
}
func receiver(c chan int) { // 异步~接受一个channel
	ret, ok := <-c // 将通道的值~赋值给ret
	fmt.Println("接受成功~", ret, ok)
}
*/

/* 练习如何关闭通道
func main() {
	c := make(chan int, 5)
	go func() { // 开启一个协程
		for i := 1; i <= 10; i++ {
			c <- i // 往c通道中存i的值
		}
		close(c) // 循环完毕~关闭通道
	}()
	for {
		if data, ok := <-c; ok { // 获取通道中的值
			fmt.Println(data, ok)
		} else {
			break
		}
	}
	fmt.Println("main结束")
}
*/
// channel 练习
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {       // 判断通道是否关闭：方式一(判断ok是否为0值)
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println(i) // 判断通道是否关闭: 方式二(range结束就是关闭，通常用range方式判断)
	}
}
