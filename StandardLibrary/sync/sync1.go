package main

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

// sync: 提供同步单元
// 一般go语言中使用sync.WaitGroup来实现并发任务的同步
// a
//Add(delta int) +delta的计数器
// d
//Done() 计数器减1
// Wait() 阻塞直到计数器为0

// type Once struct {} // Once 是只执行一次的动作对象
var DoOnce sync.Once
var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	//TestOnce()
	//TestBuildInMap()
	//TestSyncMap()
	TestAtomicBehavior()
}

func TestOnce() {
	once := func() {
		fmt.Println("Only do once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		DoOnce.Do(once)
		fmt.Println("goog")
		done <- true
	}
	for i := range done {
		a := <-done
		fmt.Println(a, i)
	}
	time.Sleep(time.Second)
}

// TestBuildInMap 测试go内置map并发行为，go内置map不带有并发行为，会报错。需要附加一个互斥锁
func TestBuildInMap() {
	i, m := 0, make(map[string]int)
	for ; i < 20; i++ {
		wg.Add(1)
		i := i
		go func() {
			mutex.Lock()
			m[strconv.Itoa(i)] = i
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(m)
}

func TestSyncMap() {
	sm := sync.Map{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		i := i
		go func() {
			sm.Store(strconv.Itoa(i), i)
			wg.Done()
		}()
	}
	wg.Wait()
	sm.Range(func(key, value interface{}) bool {
		fmt.Println("key:", key, "value:", value)
		return true
	})
}

// atomic: 对于基础数据~采用原子化操作，性能更加高

func TestAtomicBehavior() {
	var a int32
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&a, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(a)
}
