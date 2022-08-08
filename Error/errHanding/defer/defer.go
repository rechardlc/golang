package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

// userError 定义一个userError接口，其中嵌套error接口，定义Message方法
type userError interface {
	error
	Message() string
}

const prefix = "/list/"

// uErr 定义uErr类型，并且实现Message、Error方法，Message、Error都定义方法，说明实现了userError接口
type uErr string

func (e uErr) Message() string {
	return string(e)
}
func (e uErr) Error() string {
	return e.Message()
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
func writeFile(filename string) {
	//file, err := os.Create(filename) // os.Create创建文件~返回一个open文件
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	err = errors.New("this is a custom error")
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			// 最简单的出错处理
			fmt.Println("origin error: ", pathError.Op, pathError.Path, pathError.Err)
		}
		fmt.Println("filename is exist:", err.Error())
		return
	}
	defer file.Close()              // defer表现为栈形式，后进先出
	writer := bufio.NewWriter(file) // 构建一个缓存
	defer writer.Flush()            // 将缓存写入
	f := fibonacci()
	for i := 0; i <= 20; i++ {
		fmt.Fprintln(writer, "第", i, "项是", f()) // 写入值
	}
}
func handleFileList(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path
	fmt.Println("path:", path)
	fmt.Println("index:", strings.Index(path, prefix))
	if strings.Index(path, prefix) != 0 {
		return uErr("path must start with " + prefix)
	}
	path = path[len(prefix):]
	fmt.Println("after path: ", path)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}

/*
	统一出错处理
	函数式编程思想~
*/
func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		code := http.StatusOK
		defer func() {
			r := recover()
			if err, ok := r.(error); ok {
				code = http.StatusInternalServerError
				http.Error(writer, http.StatusText(code), code)
			} else if err != nil {
				panic(err)
			}
		}()
		err := handler(writer, request)
		// 通过uErr类型实现userError接口，这里就可以使用.(type)的方式进行类型断言
		if userErr, ok := err.(userError); ok {
			http.Error(writer, userErr.Message(), http.StatusBadRequest)
			return
		}
		if err != nil {
			log.Print(err.Error())
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}
func webServer() {
	http.HandleFunc("/", errWrapper(handleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}

}
func main() {
	webServer()
	//writeFile("fib.txt")
}
