package main

import (
	"flag"
	"learngo/crawler_distributed/rpcsupport"
	"fmt"
	"learngo/crawler_distributed/worker"
)

//go语言命令行参数：使用flag库
//第一个参数是命令行参数名，比如下面例子就应该输入--port，第二个参数是默认值，第三个参数是使用帮助，输入--help会有帮助
var rpcpPort = flag.Int("port",0,"input the rpc server port that you will listen,Usage:--port=")
func main() {
	flag.Parse()
	if *rpcpPort == 0 {
		fmt.Println("Must specify a rpc server port")
		return
	}
	//err := rpcsupport.ServeRpc(fmt.Sprintf(":%d",config.WorkerPort0),worker.CrawlService{})
	err := rpcsupport.ServeRpc(fmt.Sprintf(":%d",*rpcpPort),worker.CrawlService{})
	if err != nil {
		panic(err)
	}
}
