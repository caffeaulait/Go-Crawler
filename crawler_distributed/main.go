package main

import (
	"flag"
	"go_crawler/crawler/engine"
	"go_crawler/crawler/parser"
	"go_crawler/crawler/scheduler"
	"go_crawler/crawler_distributed/config"
	itemsaver "go_crawler/crawler_distributed/persist/client"
	"go_crawler/crawler_distributed/rpcsupport"
	worker "go_crawler/crawler_distributed/worker/client"
	"log"
	"net/rpc"
	"strings"
)


var (
	itemSaverHost = flag.String("itemsaver", "", "item saver host")
	workHosts = flag.String("workers", "", "work hosts(comma separated)")
)

func main()  {
	flag.Parse()
	itemChan, err := itemsaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}
	pool := createClientPool(strings.Split(*workHosts, ","))
	processor := worker.CreateProcessor(pool)
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 20,
		ItemChan: itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url: "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
}

func createClientPool(hosts []string) chan *rpc.Client{
	var clients []*rpc.Client
	for _, host := range hosts {
		client, err := rpcsupport.NewClient(host)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", host)
		} else {
			log.Printf("error connecting to %s: %v", host, err)
		}
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}