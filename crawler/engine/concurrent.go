package engine

import (
	"log"
)

type Processor func(Request) (ParseResult, error)

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan Item
	RequestProcessor Processor
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	client := CreateRedisClient()
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		go e.createWorker(e.Scheduler.GetWorker(), out, e.Scheduler)
	}
	for _, r := range seeds {
		if IsDuplicate(client, r.Url) {
			log.Printf("Duplicate request: %s", r.Url)
			continue
		}
		e.Scheduler.Submit(r)
	}
	for {
		result := <-out
		for _, item := range result.Items {
			go func(i Item) { e.ItemChan <- i }(item)
		}
		for _, request := range result.Requests {
			if IsDuplicate(client, request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	for {
		ready.WorkerReady(in)
		request := <-in
		result, err := e.RequestProcessor(request)
		if err != nil {
			continue
		}
		out <- result
	}
}

//var visitedUrls = make(map[string]bool)
//
//func isDuplicate(url string) bool {
//	if visitedUrls[url] {
//		return true
//	}
//	visitedUrls[url] = true
//	return false
//}
