package engine

import (
	"fetcher"
	"fmt"
)

type SimpleEngine struct {

}
func (se SimpleEngine) Run(seeds ...Request)  {
	// Request 队列
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for  len(requests) > 0  {
		r := requests[0]
		requests = requests[1:]

		//log.Info("[Fetch] url: %s", r.Url)
		parseResult, err := worker(r) // <--修改了该内容
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			//log.Info("Got item %s", item)
			fmt.Printf("got item %v\n", item)
		}
	}
}

// worker --> 添加了该内容
func worker(r Request) (ParseResult,error) {
	fmt.Printf("[Fetch] url : %v\n", r.Url)
	//body, err := fetcher.Fetch(r.Url)
	body, err := fetcher.FetchUrl(r.Url)
	if err != nil {
		//log.Error("[Fetcher error]：fetcher url %s:%v", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body), nil
}