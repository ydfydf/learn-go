package main

import (
	"testing"
	"learngo/crawler_distributed/rpcsupport"
	"time"
	"learngo/crawler_distributed/config"
	"learngo/crawler_distributed/worker"
)

func TestCrawlService(t *testing.T) {
	const host = ":9001"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})

	time.Sleep(time.Second)

	clinet , err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	req := worker.Request{
		Url: "http://album.zhenai.com/u/1207773016",
		Parser: worker.SerializedParser{
			FunctionName: config.ParseProfile,
			Args: "无悔",
		},
	}
	var result worker.ParseResult
	err = clinet.Call(config.CrawServiceRpc,req,&result)
	if err != nil {
		t.Error(err)
	}else {
		t.Log(result)
	}
}
