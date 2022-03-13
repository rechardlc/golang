package main

import (
	"fmt"
	"time"
)

/*
	time.NewTimer() // 新创建一个timer，返回C通道
	time.Sleep() // 阻塞go程序运行时间
*/
func main() {
	timer1 := time.NewTimer(time.Second)
	fmt.Println(timer1)
	t1 := time.Now()
	fmt.Println(t1)
	t2 := <-timer1.C
	fmt.Println(t2)
	//for {
	//	<-timer1.C
	//	fmt.Println("时间到~")
	//}
	fmt.Println("5秒开始")
	t3 := time.After(5 * time.Second)
	fmt.Println("5秒到", &t3)

	ticker := time.NewTicker(1 * time.Second)
	i := 0
	go func() {
		for {
			i++
			fmt.Println(<-ticker.C)
			if i == 5 {
				ticker.Stop()
				break
			}
		}
	}()
	for {

	}
}
