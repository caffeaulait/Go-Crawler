package client

import (
	"go_crawler/crawler/engine"
	"go_crawler/crawler_distributed/config"
	"go_crawler/crawler_distributed/worker"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	return func(request engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(request)
		var sResult worker.ParseResult
		c := <- clientChan
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeSerializeResult(sResult)
	}
}