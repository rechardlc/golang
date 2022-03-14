Golang函数调动的机制
1. 函数调用是值拷贝:基本数据类型和数组都是值传递
   * 值拷贝结果如下: 指针地址完全不一样
       ```go
       package main
       import "fmt"
       func main() {
           n1 := 10
           s1 := make([]byte, 0)
           fmt.Printf("n1: %p\n", &n1) // n1地址: 0xc0000b0008
           fmt.Printf("s1: %p\n", &s1) // s1地址:  0xc00000c030
           func(n2 int, s2 []byte) {
               fmt.Printf("n2: %p\n", &n2) // n2地址: 0xc0000b0018
               fmt.Printf("s2: %p\n", &s2) // s2地址: 0xc00000c048
           }(n1, s1)
       }
       ```