package parser

import (
	"io/ioutil"
	"testing"
)

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
	result := ParseCityList(contents)
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Requests))
	}
}
