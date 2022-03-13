package main

import (
	"errors"
	"fmt"
)

func main() {
	test()
	channel()
	protectNextToDone(1, 2)
	checkUseStandardLibraryErrors()
	checkTry()
}

// test recover 与 panic一起使用，recover只有在defer调用函数中才会有效，defer必须在panic之前
func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("panic error!")
}
func channel() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	ch := make(chan int, 8)
	close(ch)
	ch <- -1
}

func protectNextToDone(x, y int) {
	var z int
	// 使用闭包，使后面的函数继续执行~panic会中断函数执行
	func() {
		defer func() {
			if err := recover(); err != nil {
				z = 0
				fmt.Println(err)
			}
		}()
		panic("test panic")
		return
	}()
	fmt.Printf("x / y = %d\n", z)
}

var ErrDivByZero = errors.New("division by zero")

func useStandardLibraryErrors(x, y int) (int, error) {
	if y == 0 {
		return 0, ErrDivByZero
	}
	return x / y, nil
}
func checkUseStandardLibraryErrors() {
	defer func() {
		fmt.Println(recover())
	}()
	switch z, err := useStandardLibraryErrors(10, 1); err {
	case nil:
		fmt.Println(z)
	case ErrDivByZero:
		panic(err)
	}
}
func Try(fun func(), catch func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			catch(err)
		}
	}()
	fun()
}
func checkTry() {
	Try(func() {
		panic("check panic")
	}, func(err interface{}) {
		fmt.Println(err)
	})
}
