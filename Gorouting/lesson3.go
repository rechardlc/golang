/*
	Goroutine
	1. 轻量级"线程"
	2. 非抢占式多任务处理，由协程主动交出控制权
	3. 编译器、解析器、虚拟机层面的多任务
	4. 多个协程可能在一个或多个线程上运行

	Goroutine可能交出控制权的方式
	1. I/O，select
	2. channel
	3. 等待锁(sync库)
	4. 函数调用(可能)
	5. runtime.Gosched()方法
	上述方式只是参考，并不能保证切换
*/
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i] += 2
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
