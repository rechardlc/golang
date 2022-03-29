# 变量声明
1. var 或者 := 声明变量
2. 定义的每一个变量都有对应的一个初始值：如int、float32等默认值为0，bool默认值为false、string默认值为""
## 定义变量的方式
```go
// 直接声明
var a int 
var b string
// 集中定义
var (
	c int 
	d string
)
// 赋值行为
var e, f = false, uint8
// 或者在函数内使用
g, h := 0, true
```
## 注意
- var可在全局使用，:= 方式只能在函数中使用 
- 能简写，尽量简写

# 常量定义
1. 常量用const去定义
2. const定义的方式与var一样

## const实现其他语言的枚举类型方式
### 普通枚举方式
```go
const (
	a = 0
	b = 1
	c = 2
)
```
### 自增枚举方式
- Go通过iota关键字方式实现自增枚举
```go
// iota从0开始自增
const 
	a iota // 0
	b // 1
	c // 2
)
```
- Go通过iota字段自定义枚举
```go
const (
	a = 1
	b = iota // b = 1: iota 在第二位置，所以b = 1, iota从0位置开始自增
	c // c = 1
)
```
```go
// 可以通过-跳过iota// iota可以使用计算公司
const (
	a = iota + 2 //a = 2
	- // 跳过
    b // b = 4
)
```

## 注意
- Go语言的大小存在特殊含义，不能像其他语言一样，将const定义的类型为大写字母

# 类型转化
1. Go语言只有强制类型转化，没有隐式转化方式
```go
a, b := 3, 5
var c int
// 如下：math.Sqrt接受一个float64类型的参数，所以需要调用float64()只能通过显式的去转化，
// 最后c为一个int类型，所以需要通过int()去显式的将float64转化为int类型，最后赋值给c
c = int(math.Sqrt(float64(a*a + b*b)))
```