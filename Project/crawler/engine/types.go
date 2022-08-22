package engine

import (
	"fmt"
)

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func userParser(contents []byte) ParseResult {
	fmt.Printf("%s\n", contents)
	return ParseResult{}
}

func NilParser(contents []byte) ParseResult {
	//var result ParseResult
	//personUrlRe := regexp.MustCompile(`<a href="(http://album.zhenai.com/u/\w+)" target="_blank">[^<]+</a>`)
	//matches := personUrlRe.FindAllSubmatch(contents, -1)
	//for _, m := range matches {
	//	result.Requests = append(result.Requests, Request{
	//		Url:        string(m[1]),
	//		ParserFunc: userParser,
	//	})
	//}
	//return result
	return ParseResult{}
}
