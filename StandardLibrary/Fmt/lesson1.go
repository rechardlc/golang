package main

import "fmt"

// Print系列函数：Print系列函数会将内容输出到系统的标准输出，
// 区别在于Print函数直接输出内容，Printf函数支持格式化输出字符串，
// Println函数会在输出内容的结尾添加一个换行符。

//func.md main() {
//	n, err := fmt.Print("在终端打印信息。%v\n", "dell")
//	name := "dell"
//	fn, ferr := fmt.Printf("我是：%s\n", name)
//	fmt.Println("在终端打印单独一行显示", n, err, fn, ferr)
//}

// Fprint系列函数
//func.md main() {
//	_, err2 := fmt.Fprintln(os.Stdout, "向标准输出写入内容")
//	if err2 != nil {
//		return
//	}
//	fileObj, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
//	if err != nil {
//		fmt.Println("打开文件出错, err", err)
//	}
//	name := "richard"
//	_, err3 := fmt.Fprintf(fileObj, "往文件中写入信息：%s", name)
//	if err3 != nil {
//		return
//	}
//}

// Sprint系列函数
//func.md main() {
//	var name, age = "dell", 18
//	s1 := fmt.Sprint("richard\n")
//	s2 := fmt.Sprintf("name: %s, age:%d\n", name, age)
//	fmt.Println(s1, s2)
//}

// Errorf函数，通过Errorf函数format自定义错误类型
func main() {
	err := fmt.Errorf("这个是一个错误信息~")
	fmt.Println(err)
}
