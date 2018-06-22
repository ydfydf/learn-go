package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/scheduler"
	"learngo/crawler/zhenai/parser"
	"learngo/crawler_distributed/persist/client"
	"learngo/crawler_distributed/config"
	client2 "learngo/crawler_distributed/worker/client"
	"net/rpc"
	"learngo/crawler_distributed/rpcsupport"
	"log"
	"flag"
	"strings"
)

const seedurl = "http://www.zhenai.com/zhenghun"

var (
	itemSaverHost = flag.String("itemsaver_host","","itemsaver host")
	workerHosts = flag.String("worker_hosts","","worker hosts (comma separated)")
)
func main() {
	flag.Parse()
	var seed engine.Request
	seed = engine.Request{
		Url : seedurl,
		Parser : engine.NewFuncParser(parser.ParseCityList,config.ParseCityList),
	}
	itemChan , err := client.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createRpcClientPool(
		strings.Split(*workerHosts,","),
	)
	processor  := client2.CreateProcessor(pool)

	e := engine.ConcurrentEngine{//这里的e的数据类型是接口，此为接口变量赋值
		Scheduler: &scheduler.QueuedScheduler{},//Scheduler: &scheduler.SimpleScheduler{},也是接口变量赋值，确定使用哪个实现的方法
		WorkerCount: 100,
		ItemChan: itemChan, //persist.ItemSaver(),这里会先调用persist.ItemSaver()来生成一个chan interface{}
		RequestProcessor: processor,
	}
	e.Run(seed)//接口变量e调用engine.ConcurrentEngine为其实现的方法
}

func createRpcClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _ , h := range hosts {
		client , err := rpcsupport.NewClient(h)
		if err == nil{
			clients = append(clients,client)
			log.Printf("Connected to %s", h)
		}else {
			log.Printf("Error connecting to %s: %v", h,err)
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