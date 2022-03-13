package main

import (
	"fmt"
	"sync" // 提供基本同步单元
	"time"
)

// 启动多个goroutine, 多个goroutine是并发执行的，然而goroutine的调度是随机的
var wg sync.WaitGroup //用于等待一组线程的结束

func hello(i int) {
	defer wg.Done() // 每执行一个hello方法后，减少一个WaitGroup值，在函数最后执行~这个里使用defer延迟
	fmt.Println("hello Goroutine", i)
}
func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)   // 每次循环，增加一个WaitGroup值
		go hello(i) // 开启协程执行
	}
	//runtime.Gosched()
	wg.Wait() //等待WaitGroup结果，如WaitGroup为0，不再阻塞主线程执行
	fmt.Println("done")
	verification()
}

// 验证主协程退出情况,主协程退出，其他任务也会跟着一起终止
func verification() {
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine: i = %v\n", i)
			time.Sleep(time.Second)
		}
	}()
	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i = %v\n", i)
		time.Sleep(time.Second)
		if i == 10 {
			break
		}
	}
}
