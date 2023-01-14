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
	if runtime.GOOS == "windows" {
		name = "cmd"
		c = "/C"
	}
	arg = append([]string{c}, arg...)
	cmd := exec.Command(name, arg...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return result, err
	}
	if err = cmd.Start(); err != nil {
		return result, err
	}
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return result, err
	}
	if err = cmd.Wait(); err != nil && err.(*exec.ExitError).Stderr != nil {
		return result, err
	}
	result = string(bytes)
	return result, nil
}
