package main

import (
	"net/rpc"
	"learngo/rpc"
	"net"
	"log"
	"net/rpc/jsonrpc"
)

func main() {
	//使用rpc来包装Service
	err := rpc.Register(rpcdemo.DemoService{})
	if err != nil{
		panic(err)
	}
	listener, err := net.Listen("tcp",":1234")
	if err != nil{
		panic(err)
	}

	for {
		conn , err := listener.Accept()
		if err != nil {
			log.Printf("accept error : %v",err)
			continue
		}
		//使用jsonrpc来处理connection，调用Service里的Method
		go jsonrpc.ServeConn(conn)
	}
}
