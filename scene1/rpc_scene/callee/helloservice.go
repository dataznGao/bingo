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
