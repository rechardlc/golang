package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func getCircleArea(radius float32) (area float32, err error) {
	if radius < 0 {
		err = errors.New("半径不能为负数")
		return
	}
	area = 3.14 * radius * radius
	return
}
func test03() {
	_, err := getCircleArea(-4)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("这里有没有执行")
}
func test04() {
	test03()
	fmt.Println("test04 done")
}
func main() {
	test04()
	TestOpenFunc()
}

// PathError 自定义错误类型
type PathError struct {
	path, op, createTime, message string
}

func (p PathError) Error() string {
	return fmt.Sprintf("path:%s\n op=%s\n createTime=%s\n message=%s\n", p.path, p.op, p.createTime, p.message)
}
func Open(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return &PathError{
			path:       fileName,
			op:         "read",
			message:    err.Error(),
			createTime: fmt.Sprintf("%v", time.Now()),
		}
	}
	defer file.Close()
	return nil
}
func TestOpenFunc() {
	err := Open("/User/html/dell/text.txt")
	switch v := err.(type) {
	case *PathError:
		fmt.Println("get path error:\n", v)
	default:
		fmt.Println("open file done!")
	}
}
