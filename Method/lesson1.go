package main

import "fmt"

type User struct {
	Name, Email string
}
type T struct {
	int // 匿名字段~一个结构体只能有一个类型~可以通过类型访问~
}

/*
	1. 方法若不接受指针类型，无论是否传入的指针类型还是值类型，最终都会是值类型,反之接受指针类型
	2. 方法与普通函数的区别
		i: 在于与接受函数的区别~普通函数只能接受对应的值（指针）类型，方法无论接受的值（指针），都可以传递值和指针类型
*/
// User
func (u *User) Notify() {
	u.Email = "change@c.com"
	fmt.Printf("%v : %v\n", u.Name, u.Email)
}

func (t *T) TestTType() {
	fmt.Println("类型T方法集包含了全部receiver T 方法")
}

func main() {
	u1 := User{Name: "dell", Email: "golang@email.com"}
	u1.Notify()
	fmt.Println(u1.Email)
	u2 := &User{Name: "richard", Email: "richard@qq.com"}
	u2.Notify()
	fmt.Println(u2.Email)

	t1 := T{1}
	fmt.Printf("t1 is :%v\n", t1)
	fmt.Println(t1.int)
	t1.TestTType()
}
