package main

import (
	"fmt"
	"sync"
	"time"
)

// go语言一般不使用传统的同步机制，方法如：WaitGroup，Mutex，Cond，直接使用select方式去调度
// 通过command：go run -race *.go 可以查看代码是否存在资源竞争
// 一般来说，直接使用系统的原子化操作,不用自己造轮子
type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) increment() {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.value++
}
func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}
func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Second)
	fmt.Println(a.get())
}
