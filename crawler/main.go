package main

import (
	"ccmouce/crawler/engine"
	"ccmouce/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc:parser.ParseCityList,
	})
}
