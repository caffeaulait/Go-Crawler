package main

import (
	"go_crawler/crawler/config"
	"go_crawler/crawler/engine"
	"go_crawler/crawler/parser"
	"go_crawler/crawler/persist"
	"go_crawler/crawler/scheduler"
)

func main()  {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount: 20,
		ItemChan: itemChan,
		RequestProcessor: engine.Worker,
	}
	e.Run(engine.Request{
		Url: "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
}

