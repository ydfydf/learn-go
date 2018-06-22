package scheduler

import "learngo/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(w chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

//func (s *SimpleScheduler) ConfigureMasterWorkerChan(e chan engine.Request) {
//	s.workerChan = e
//}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {s.workerChan <- r}() //另开一个goroutine，让函数返回
}

