package engine

import (
	"go_crawler/crawler/fetcher"
	"log"
)

func Worker(r Request) (ParseResult, error) {
	//log.Printf("fetching %s\n", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("fail to fetch url: %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.Parser.Parse(body, r.Url), nil
}
