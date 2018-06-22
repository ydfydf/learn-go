package engine

import (
	"log"
	"learngo/crawler/fetcher"
)

func  Worker(r Request) (ParserResult,error) {
	log.Printf("Fetching URL:%s\n",r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher:error fetching url %s: %v",r.Url,err)
		return ParserResult{}, err
	}
	return  r.Parser.Parse(body,r.Url),nil
}
