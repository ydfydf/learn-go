package rpcsupport

import (
	"net/rpc"
	"net"
	"log"
	"net/rpc/jsonrpc"
)

func ServeRpc(host string,service interface{}) error{
	err := rpc.Register(service)
	if err != nil{
		return err
	}
	listener, err := net.Listen("tcp",host)
	if err != nil{
		return err
	}
	log.Printf("Listening the rpc server port on %s",host)

	for {
		conn , err := listener.Accept()
		if err != nil {
			log.Printf("accept error : %v",err)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
	return nil
}

func NewClient(host string) (*rpc.Client, error) {
	conn ,err := net.Dial("tcp",host)
	if err != nil {
		return nil,err
	}
	//使用jsonrpc来包装这个connection
	return jsonrpc.NewClient(conn),nil

}