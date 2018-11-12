package main

import (
	"engine"
	"scheduler"
	"zhengai/parse"
)

func main() {
	/*e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount:100,
	}*/
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkerCount:100,
	}
	/*
	// 获取全部数据
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parse.ParserCityList,
	})*/

	// 获取上海的数据
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parse.ParserCity,
	})
}