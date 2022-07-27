// go语言的duck typing: 具有py、c++的灵活性，也有java的类型定义
// go语言的由使用者来定义接受者的
package main

import "fmt"

/**
go语言的接口定义与实现：由使用者来定义接受者的
如下例子：
1. 定义MockRetriever结构体
2. 实现MockRetriever其中的Get方法
3. 使用者定义接口,如定义Retriever接口，如需要使用get方法，就去定义get方法，不使用也就可以不定义
4. 定义download方法，将Retriever接口定义传递进去
5. 调用download方法
*/

// MockRetriever
/* ======================= MockRetriever结构体定义及其Get方法 ================= */
type MockRetriever struct {
	Contents string
}

func (r MockRetriever) Get(url string) string {
	return r.Contents
}

/* ======================= 定义结束 ============== */

type Retriever interface {
	Get(url string) string
}
type Poster interface {
	Post(url string, form map[string]string) string
}

// RetrieverPoster 接口的组合
type RetrieverPoster interface {
	Poster
	Retriever
}

func download(r Retriever) string {
	return r.Get("www.john.cn")
}
func post(poster Poster) {
	poster.Post("www.baidu.com", map[string]string{
		"name": "dell",
	})
}

// 接口组合后~使用
func session(s RetrieverPoster) {
	s.Get("")
	s.Post("", map[string]string{})
}
func main() {
	var result Retriever
	result = &MockRetriever{Contents: "https://www.john.cn"}
	fmt.Printf("%T", result)
	mock, ok := result.(*MockRetriever) // T.(type)
	fmt.Println(mock, ok)
	switch v := result.(type) { // 判断类型方式: T.(type)
	case MockRetriever:
		fmt.Println(result, v.Contents)
	}
	mockResult := download(result)
	fmt.Println(mockResult)
	pointType(1)
}

// Queue Any interface {} 定义任意类型
type Queue []interface{}

// 对于通用类型，进行强制类型转化，用.(type)的方式处理
func pointType(a int) int {
	var q Queue
	q = append(q, a)  // 如取Queue类型，只要注入int类型，直接写成a.(int)方式,只有在运行时才会报错
	return q[0].(int) // 返回int类型，取q[0].(int),只有在运行时才会报错
}
