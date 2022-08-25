package engine

import (
	"crawler/fetcher"
	"log"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	// 值传递，可以直接修改seeds
	count := 0
	for len(seeds) > 0 {
		req := seeds[0] // 出列，广度搜索优先算法
		seeds = seeds[1:]
		parseResult, err := worker(req)
		if err != nil {
			continue
		}
		seeds = append(seeds, parseResult.Requests...)
		for _, item := range parseResult.Items {
			count++
			//log.Printf("Got item %#v\n", item)
			log.Printf("current Item: #%d;\nGot item: %v\n", count, item)
		}
	}
}
func worker(req Request) (ParseResult, error) {
	log.Printf("Fetching %s\n", req.Url)
	body, err := fetcher.Fetch(req.Url) // 先执行fetch方法
	if err != nil {
		log.Printf("fetcher: error fetching url %s; 异常: %v\n", req.Url, err)
		return ParseResult{}, err
	}
	return req.ParserFunc(body), nil
}
