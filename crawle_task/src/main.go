package main

import (
	"engine"
	"zhengai/parse"
)

func main() {
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parse.ParserCityList,
	})
}