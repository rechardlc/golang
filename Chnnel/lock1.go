package main

import (
	"fmt"
	"sync"
	"time"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex // 互斥锁：多个goroutine修改共享资源，
var rwlock sync.RWMutex

// 添加互斥锁，会有且仅有一个goroutine进入临界区，其他goroutine则在等待状态
// 待互斥锁释放后，等待的才能进去，被等待的goroutine是随机的

func main() {
	TestRwLock()
	//wg.Add(2)
	//go add() // 同时开始两个协程，去修改同一个x(共享资源)的值，存在生态竞争
	//go add()
	//wg.Wait()
	//fmt.Println(x)
}
func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock() // 互斥锁
		x = x + 1   //
		lock.Unlock()
	}
	wg.Done()
}

// 读写锁测试

func write() {
	rwlock.Lock() // 创建写锁
	x = x + 1
	time.Sleep(10 * time.Second)
	rwlock.Unlock() // 接触写锁
	wg.Done()
}
func read() {
	rwlock.RLock() // 创建读锁
	time.Sleep(time.Millisecond)
	rwlock.RUnlock() // 结束读锁
	wg.Done()
}
func TestRwLock() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
