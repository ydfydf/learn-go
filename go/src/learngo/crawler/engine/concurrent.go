package engine

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan    chan Item
	RequestProcessor Processor
}

type Processor func(Request) (ParserResult, error)

type Scheduler interface {
	//重构，将WorkerReady提出去，放到ReadyNotifier里
	ReadyNotifier
	Submit(Request)
	//重构worker的channel，由shcduler来决定这个channel是公用还是每个worker一人一个
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(w chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request){
	out := make(chan ParserResult)
	e.Scheduler.Run() //e.Scheduler.Scheduler.requestChan和e.Scheduler.Scheduler.workerChan等待接收相应的数据，
	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker2(e.Scheduler.WorkerChan(),out,e.Scheduler)//向e.Scheduler.Scheduler.workerChan添加数据(添加channel)
	}
	for _, r := range seeds {
		e.Scheduler.Submit(r)//向e.Scheduler.Scheduler.requestChan发送Request数据
	}

	//itemCount := 0
	//for {
	//	result := <- out
	//	for _ ,item := range result.Items {
	//		log.Printf("Got item #%d: %v",itemCount,item)
	//		itemCount ++
	//	}
	//	for _, request := range result.Requests {
	//		e.Scheduler.Submit(request)
	//	}
	//}
	for {
		result := <- out
		for _ ,item := range result.Items {
			go func() {e.ItemChan <- item}() //采用开goroutine的方式存储item，假设有10w个item，要开10w个这么简单的goroutine，性能损耗是很小的
		}
		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}
//注意这里会造成死循环，在createWorker开goroutine一直在接收channle in的数据（共100个），每个goroutine执行完后发送数据给out，但是result := <- out在接收数据后又执行e.Scheduler.Submit(request)
//，又往in中添加数据，但是由于现在全部的goroutine都在工作，没有goroutine去接收in的数据，导致程序在往in中添加数据时等待，e.Scheduler.Submit(request)不能返回，导致程序造成死循环，
//所以在Submit中开goroutine，go func() {s.workerChan <- r}()，使函数不用等待就返回，这样死循环就解开了
func createWorker(in chan Request, out chan ParserResult){
	go func() {
		for {
			request := <-in
			result ,err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
func createWorker1(out chan ParserResult, s Scheduler){
	in := make(chan Request) //每个Worker创建一个自己的channel->in
	go func() {
		for {
			//tell scheduler i'm ready
			s.WorkerReady(in) //将空的in放入e.Scheduler.Scheduler.workerChan里，在e.Scheduler.Run()里会将e.Scheduler.Scheduler.workerChan里的in放入workerQ的队列里
			request := <-in //被放入workerQ队列里的in，等待e.Scheduler.Run()函数的goroutine将requestQ里的Request取出来放入workerQ里的in里，这样request := <-in就会被激活
			result ,err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func (e *ConcurrentEngine)createWorker2(in chan Request,out chan ParserResult,ready ReadyNotifier){
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result ,err := e.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}