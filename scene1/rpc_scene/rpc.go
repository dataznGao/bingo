package rpc_scene

import (
	"fmt"
	"log"
	"net/rpc"
	"strconv"
	"sync"
	"time"
)

func callRPC(cli *rpc.Client, s string) {
	var reply string
	err := cli.Call("HelloService.Hello", s, &reply)
	if err == nil {
		panic(err)
	}
	fmt.Println(reply)
}
func Motivate() {
	start := time.Now()
	cli, err := rpc.Dial("tcp", ":8080")
	if err != nil {
		log.Println("client create failure")
	}
	group := &sync.WaitGroup{}
	group.Add(10)
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			callRPC(cli, strconv.Itoa(i))
			group.Done()
		}()
	}
	group.Wait()
	end := time.Now()
	cost := end.Second() - start.Second()
	log.Printf("[INFO] cost time is %vs", cost)
	if cost > 3 {
		log.Print("[ERROR] timeout, SLA has broken!!!")
	}
}
