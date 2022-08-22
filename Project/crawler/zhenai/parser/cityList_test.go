package parser

import (
	"io/ioutil"
	"testing"
)

// 写测试用例，尽量保证被测试对象处于一个本地化状态，例如如下测试：直接通过Get的方式拉取网站，被拉取的网站若存在丢失、迁移，或者测试机器没有联网，都会存在测试失败
func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("cityList_test.html")
	if err != nil {
		panic(err)
	}
	// 将http.Get的数据读入cityList_test.html，用于本地做test
	//err = ioutil.WriteFile("cityList_test.html", contents, 0644)
	//if err != nil {
	//	panic(err)
	//}
	const resultSize = 470
	result := CityList(contents)
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Requests))
	}
	if len(result.Items) != resultSize {
		t.Errorf("result should have %d Items; but had %d", resultSize, len(result.Items))
	}
}
