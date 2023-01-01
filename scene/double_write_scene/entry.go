package double_write_scene

import (
	"encoding/json"
	"fmt"
	"github.com/dataznGao/go_drill/scene/double_write_scene/dal"
	"log"
	"net/rpc"
)

type CreateTestZnRsp struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Data       string `json:"data"`
}

// DoubleWrite 双写场景
func DoubleWrite(t *dal.TestEn) *CreateTestZnRsp {
	err := dal.AddTest(t)
	if err != nil {
		marshal, _ := json.Marshal(t)
		log.Printf("[ERROR] write to DataBase error! err: %v, data is %v\n", err, string(marshal))
		return &CreateTestZnRsp{
			Message:    "[DoubleWrite] write to database failure",
			StatusCode: 1,
			Data:       "",
		}
	}
	marshal, _ := json.Marshal(t)
	cli, err := rpc.Dial("tcp", ":8080")
	if err != nil {
		log.Printf("[ERROR] client create failure, err: %v, data is %v\n", err, string(marshal))
		err = dal.DelTest(t)
		if err != nil {
			log.Printf("[ERROR] del data failure, err: %v, data is %v\n", err, string(marshal))
		}
		return &CreateTestZnRsp{
			Message:    "[DoubleWrite] call rpc failure",
			StatusCode: 1,
			Data:       "",
		}
	}
	err = callRPC(cli, "10")
	if err != nil {
		log.Printf("[ERROR] call rpc failure, err: %v, data is %v\n", err, string(marshal))
		err = dal.DelTest(t)
		if err != nil {
			log.Printf("[ERROR] del data failure, err: %v, data is %v\n", err, string(marshal))
		}
		return &CreateTestZnRsp{
			Message:    "[DoubleWrite] call rpc failure",
			StatusCode: 1,
			Data:       "",
		}
	}
	return &CreateTestZnRsp{
		Message:    "[DoubleWrite] success",
		StatusCode: 0,
		Data:       string(marshal),
	}
}

func callRPC(cli *rpc.Client, s string) error {
	var reply string
	//在调用client.Call时，
	//第一个参数是用点号链接的RPC服务名字和方法名字，
	//第二和第三个参数分别我们定义RPC方法的两个参数。
	err := cli.Call("HelloService.Hello", s, &reply)
	if err != nil {
		return err
	}
	fmt.Println(reply)
	return nil
}
