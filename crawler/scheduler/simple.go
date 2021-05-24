package scheduler

import "go_crawler/crawler/engine"

type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {s.WorkerChan <- r}()
}

func (s *SimpleScheduler) GetWorker() chan engine.Request{
	return s.WorkerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {

}

func (s *SimpleScheduler) Run() {
	s.WorkerChan = make(chan engine.Request)
}