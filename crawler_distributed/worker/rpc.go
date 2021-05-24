package worker

import "go_crawler/crawler/engine"

type CrawlService struct {

}

func (*CrawlService) Process(request Request, result *ParseResult) error {
	engineReq, err := DeSerializeRequest(request)
	if err != nil {
		return err
	}
	engineResult, err := engine.Worker(engineReq)
	*result = SerializeResult(engineResult)
	return nil
}
