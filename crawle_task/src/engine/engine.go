package engine

import (
	"fetcher"
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
)

func Run(seeds ...Request)  {
	// Request 队列
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for  len(requests) > 0  {
		r := requests[0]
		requests = requests[1:]

		//log.Info("[Fetch] url: %s", r.Url)
		fmt.Printf("[Fetch] url : %v\n", r.Url)
		//body, err := fetcher.Fetch(r.Url)
		body, err := fetcher.FetchUrl(r.Url)
		if err != nil {
			log.Error("[Fetcher error]：fetcher url %s:%v", r.Url, err)
			continue
		}

		parseResult := r.ParserFunc(body)
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			//log.Info("Got item %s", item)
			fmt.Printf("got item %v\n", item)
		}
	}
}
