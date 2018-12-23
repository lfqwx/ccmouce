package engine

import (
	"ccmouce/crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _,r:=range seeds{
		requests=append(requests,r)
	}
	for len(requests)>0{
		r:=requests[0]
		requests=requests[1:]
		parseResult,err:=worker(r)
		if err!=nil{
			continue
		}
		requests=append(requests,parseResult.Requests...)
		for _,item:=range parseResult.Iterms {
			log.Printf("Got item %v",item)
		}

	}
}

//建立并发worker
func worker(r Request)  (ParseResult,error){
	log.Printf("Fetching %s",r.Url)
	body,err:=fetcher.Fetch(r.Url)
	if err!=nil{
		log.Printf("Fetcher:error",r.Url)
		return ParseResult{},err
	}
	return r.ParserFunc(body),nil
}
