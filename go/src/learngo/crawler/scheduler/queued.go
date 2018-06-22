package scheduler

import "learngo/crawler/engine"

//队列的分发器scheduler

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s QueuedScheduler) WorkerReady(w chan engine.Request){
	s.workerChan <- w
}


func (s *QueuedScheduler) Run() {
	//此函数作用为：将收到的Request放入requestQ队列里去排队，将chan engine.Request放入workerQ进行排队，然后将requestQ里的Request导入
	//workerQ里的chan engine.Request，这样在s.workerChan里的in就有数据了(s.workerChan里的chan engine.Request都是在createWorker1里创建的空channel)，
	//createWorker1里的request := <-in就被激活了
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ  []chan engine.Request
		for{
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <- s.requestChan:
				requestQ = append(requestQ,r)
			case w := <- s.workerChan:
				workerQ = append(workerQ,w)
			case activeWorker <- activeRequest:
					requestQ = requestQ[1:]
					workerQ = workerQ[1:]
			}
		}
	}()
}

