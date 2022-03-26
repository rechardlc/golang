# 变量声明
1. var 或者 := 声明变量
2. var能在任何地方声明，:=只能在函数中使用
3. 定义的每一个变量都有对应的一个初始值：如int、float32等默认值为0，bool默认值为false、string默认值为""
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
<font face="黑体">我是黑体字</font>