package double_write_scene

import (
	"encoding/json"
	"fmt"
	"github.com/dataznGao/bingo/scene/double_write_scene/dal"
	"log"
	"net/rpc"
)

type CreateTestZnRsp struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Data       string `json:"data"`
}

func DoubleWrite(t *dal.TestEn) *CreateTestZnRsp {
	err := dal.AddTest(t)
	if err != nil {
		marshal, _ := json.Marshal(t)
		log.Printf("[ERROR] write to DataBase error! err: %v, data is %v\n", err, string(marshal))
		return &CreateTestZnRsp{Message: "[DoubleWrite] write to database failure", StatusCode: 1, Data: ""}
	}
	marshal, _ := json.Marshal(t)
	cli, err := rpc.Dial("tcp", ":8080")
	if err != nil {
		log.Printf("[ERROR] client create failure, err: %v, data is %v\n", err, string(marshal))
		err = dal.DelTest(t)
		if err != nil {
			log.Printf("[ERROR] del data failure, err: %v, data is %v\n", err, string(marshal))
		}
		return &CreateTestZnRsp{Message: "[DoubleWrite] call rpc failure", StatusCode: 1, Data: ""}
	}
	err = callRPC(cli, "10")
	if err != nil {
		log.Printf("[ERROR] call rpc failure, err: %v, data is %v\n", err, string(marshal))
		err = dal.DelTest(t)
		if err != nil {
			log.Printf("[ERROR] del data failure, err: %v, data is %v\n", err, string(marshal))
		}
		return &CreateTestZnRsp{Message: "[DoubleWrite] call rpc failure", StatusCode: 1, Data: ""}
	}
	return &CreateTestZnRsp{Message: "[DoubleWrite] success", StatusCode: 0, Data: string(marshal)}
}
func callRPC(cli *rpc.Client, s string) error {
	var reply string
	err := cli.Call("HelloService.Hello", s, &reply)
	if err != nil {
		return err
	}
	fmt.Println(reply)
	return nil
}
