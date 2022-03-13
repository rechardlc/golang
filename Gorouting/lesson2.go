package main

import (
	"fmt"
	"runtime"
	"time"
)

// runtime包 runtime.Gosched() // 事还是要做
//func main() {
//	go func(s string) {
//		for i := 0; i < 2; i++ {
//			fmt.Println(s)
//		}
//	}("world")
//	for i := 0; i < 2; i++ {
//		runtime.Gosched()
//		fmt.Println("hello")
//	}
//}

//
func main() {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			runtime.Goexit()
			defer fmt.Println("C,defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	for {
		fmt.Println("主线程~执行")
		time.Sleep(time.Second * 60)
		break
	}
}
