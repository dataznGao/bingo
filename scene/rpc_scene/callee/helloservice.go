package main

import (
	"net"
	"net/rpc"
	"time"
)

type HelloService struct{}

func (h *HelloService) Hello(request string, reply *string) error {
	time.Sleep(time.Millisecond * 500)
	*reply = "hello:" + request
	return nil
}

func main() {
	err := rpc.RegisterName("HelloService", new(HelloService))
	if err != nil {
		panic(err)
	}

	//然后我们建立一个唯一的TCP链接，并且通过rpc.ServeConn函数在该TCP链接上为对方提供RPC服务。
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	con, err := lis.Accept()
	if err != nil {
		panic(err)
	}

	rpc.ServeConn(con)
}
