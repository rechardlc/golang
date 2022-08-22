package parser

import (
	"crawler/engine"
	"regexp"
)

var cityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/\w+)"[^>]*>([^<]+)</a>`)

// CityList /* 从入口进来后，进入第二个解析器，解析城市列表
func CityList(contents []byte) (result engine.ParseResult) {
	matches := cityListRe.FindAllSubmatch(contents, -1)
	limit := 2
	for _, m := range matches {
		result.Items = append(result.Items, "City: "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: City,
		})
		limit--
		if limit == 0 {
			break
		}
	}
	return result
}
