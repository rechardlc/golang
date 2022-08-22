package engine

import (
	"crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	// 值传递，可以直接修改seeds
	for len(seeds) > 0 {
		r := seeds[0] // 出列
		seeds = seeds[1:]
		log.Printf("Fetching %s\n", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("fetcher: error fetching url %s: %v\n", r.Url, err)
			continue
		}
		parseResult := r.ParserFunc(body)
		seeds = append(seeds, parseResult.Requests...)
		for _, item := range parseResult.Items {
			//log.Printf("Got item %#v\n", item)
			log.Printf("Got item %v\n", item)
		}
	}
}
