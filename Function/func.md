# 函数func
1. 函数接受值与var定义类似
```go
func calc(a, b int, operation string) {
	switch operation {
	    case "-": a - b
		case "*": a * b
    }
}
```
2. 可以返回多个值,一般来说func的第二个值用于返回错误
3. 返回值可以提前声明，但是不建议这样子声明，一旦函数体过于长，容易导致变量混乱，不易于维护
```go
func manyReturns()(q, e, c int) {
	q ++
	e ++
	c ++
	return
}
```
4. 函数可以作为参数
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
2. 对于匿名函数，声明的变量的用法~可以在函数构建函数~
   ```go
    package main
    import "fmt"
    func main() {
        var a = func(a, b int) {
        fmt.Println("将匿名函数~放置在main中")
     }
        // 在main函数中构建一个匿名函数，并且在main函数调用
        a(1,2)
   }
    ```
3. 值传递与指针拷贝区别：
   * 指针拷贝传递效率更加高，地址拷贝只有4个字节，而值拷贝可能存在很大的值