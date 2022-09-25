package util

import (
	_ "bufio"
	_ "fmt"
	"io/ioutil"
	"os/exec"
	"runtime"
)

func Command(arg ...string) (result string, err error) {
	name := "/bin/bash"
	c := "-c"
	// 根据系统设定不同的命令name
	if runtime.GOOS == "windows" {
		name = "cmd"
		c = "/C"
	}
	arg = append([]string{c}, arg...)
	cmd := exec.Command(name, arg...)

	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return result, err
	}

	//执行命令
	if err := cmd.Start(); err != nil {
		return result, err
	}

	//读取所有输出
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return result, err
	}

	if err := cmd.Wait(); err != nil {
		return result, err
	}

	result = string(bytes)
	return
}
