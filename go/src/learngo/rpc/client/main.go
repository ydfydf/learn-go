package main

import (
	"net"
	"net/rpc/jsonrpc"
	"learngo/rpc"
	"fmt"
)

func main() {
	conn ,err := net.Dial("tcp",":1234")
	if err != nil {
		panic(err)
	}
	//使用jsonrpc来包装这个connection
	client := jsonrpc.NewClient(conn)

	var result float64
	//var args rpcdemo.Args
	//args.A = 10
	//args.B = 3
	//err = client.Call("DemoService.Div",args,&result)
	err = client.Call("DemoService.Div",rpcdemo.Args{5,2},&result)
	if err != nil {
		panic(err)
	}else {
		fmt.Println("result=",result)
	}
}
