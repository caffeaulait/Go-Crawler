package scheduler

import (
	"go_crawler/crawler/engine"
)

type QueuedScheduler struct {
	RequestChan chan engine.Request
	WorkerChan  chan chan engine.Request
}

func (q *QueuedScheduler) GetWorker() chan engine.Request{
	return make(chan engine.Request)
}

func (q *QueuedScheduler) Submit(r engine.Request) {
	q.RequestChan <- r
}

func (q *QueuedScheduler) WorkerReady(w chan engine.Request) {
	q.WorkerChan <- w
}

func (q *QueuedScheduler) Run() {
	q.WorkerChan = make(chan chan engine.Request)
	q.RequestChan = make(chan engine.Request)
	go func() {
		var requests []engine.Request
		var workers []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requests) > 0 && len(workers) > 0 {
				activeRequest = requests[0]
				activeWorker = workers[0]
			}
			select {
				case r := <-q.RequestChan:
					requests = append(requests, r)
				case w := <-q.WorkerChan:
					workers = append(workers, w)
				case activeWorker <- activeRequest:
					workers = workers[1:]
					requests = requests[1:]
			}
		}
	}()
}
