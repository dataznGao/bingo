package main

import (
	"errors"
	"math/rand"
	"net"
	"net/rpc"
	"strconv"
	"time"
)

type HelloService struct{}

func (h *HelloService) Hello(request string, reply *string) error {
	time.Sleep(time.Millisecond * 500)
	*reply = "hello:" + request
	rand.Seed(time.Now().Unix())
	ran := rand.Intn(10)
	atoi, _ := strconv.Atoi(request)
	if ran < atoi {
		return errors.New("rpc called failed, please retry after 1 min")
	}
	return nil
}

func main() {
	err := rpc.RegisterName("HelloService", new(HelloService))
	if err != nil {
		panic(err)
	}
	for {
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
}
