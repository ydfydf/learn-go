package client

import (
	"learngo/crawler/engine"
	"learngo/crawler_distributed/config"
	"learngo/crawler_distributed/worker"
	"net/rpc"
	"log"
)
//func CreateProcessor(clients []*rpc.Client) (engine.Processor,error){
//传入rpc client，rpc client会连接不同的Rpc Server，如果使用上面的slice结构的话会在CreateProcessor里面发生争抢，需要加锁等互斥动作，go语言有一种更简单的方法，使用channel
func CreateProcessor(clientChan chan *rpc.Client) engine.Processor{

	return func(request engine.Request) (engine.ParserResult,error) {
		sReq := worker.SerializeRequest(request)
		var sResult worker.ParseResult
		log.Println("engine.Processor working")
		rpcClient := <-clientChan
		err := rpcClient.Call(config.CrawServiceRpc,sReq,&sResult)
		if err != nil {
			return engine.ParserResult{},err
		}
		return worker.DeserializeResult(sResult)
	}
}
