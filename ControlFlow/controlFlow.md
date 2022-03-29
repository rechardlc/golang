# 流程控制
## 与其他语言相比
1. for,if后面的条件没有括号
```go
for {}
if true {}
```
3. if条件中可以定义变量,使用分号分割
```go
if i := 0; i == 0 {}
```
4. 没有while
5. switch不需要break,在switch中可以使用多个条件
```go
switch {
    case i := 0 : done()
	case j := 1 : done() 
}
```