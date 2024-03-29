package util

import (
	_ "bufio"
	_ "fmt"
	"io/ioutil"
	"os/exec"
	"runtime"
)

func Command(arg ...string) (string, error) {
	name := "/bin/bash"
	result := ""
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
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return result, err
	}

	//执行命令
	if err = cmd.Start(); err != nil {
		return result, err
	}

	//读取所有输出
	bytes, err := ioutil.ReadAll(stdout)
	result = string(bytes)
	errs, err := ioutil.ReadAll(stderr)
	if err != nil {
		return string(errs), err
	}

	if err = cmd.Wait(); err != nil {
		return string(errs), err
	}

	return result, nil
}

func CommandTest(arg ...string) (string, error) {
	name := "/bin/bash"
	result := ""
	c := "-c"
	// 根据系统设定不同的命令name
	if runtime.GOOS == "windows" {
		name = "cmd"
		c = "/C"
	}
	cmd := exec.Command(name, "go mod tidy")
	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return result, err
	}

	//执行命令
	if err = cmd.Start(); err != nil {
		return result, err
	}

	//读取所有输出
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return result, err
	}

	arg = append([]string{c}, arg...)
	cmd = exec.Command(name, arg...)

	//创建获取命令输出管道
	stdout, err = cmd.StdoutPipe()
	if err != nil {
		return result, err
	}

	//执行命令
	if err = cmd.Start(); err != nil {
		return result, err
	}

	//读取所有输出
	bytes, err = ioutil.ReadAll(stdout)
	if err != nil {
		return result, err
	}

	if err = cmd.Wait(); err != nil && err.(*exec.ExitError).Stderr != nil {
		return result, err
	}

	result = string(bytes)
	return CleanCode(result), nil
}
