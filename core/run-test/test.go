package run

import (
	"errors"
	"github.com/dataznGao/bingo/constant"
	"github.com/dataznGao/bingo/util"
	"log"
	"os"
	"strings"
)

func Test(testPath string, inputPath string) (string, error) {
	// cd /Users/misery/GolandProjects/bingo && go test -v -cover scene/double_write_scene/entry_test.go scene/double_write_scene/entry.go
	// 校验testPath
	if !strings.HasSuffix(testPath, "_test.go") {
		return "", errors.New("the testPath is not end with _test.go")
	}
	if !strings.HasPrefix(testPath, inputPath) {
		return "", errors.New("the testPath or inputPath set err! please check! err")
	}
	// 获取相对路径
	relativeTestPath := util.CompareAndExchange(testPath, "", inputPath+constant.Separator)
	relativeTestPath = strings.TrimSpace(relativeTestPath)
	relativeRawFile := relativeTestPath[:len(relativeTestPath)-8] + relativeTestPath[len(relativeTestPath)-3:]
	absoluteRawFile := inputPath + constant.Separator + relativeRawFile
	setRawFile := true
	if _, err := os.Open(absoluteRawFile); err != nil {
		setRawFile = false
	}
	// 组装commend
	commend := "cd " + inputPath + " && go test -v -cover " + relativeTestPath
	if setRawFile {
		commend += " " + relativeRawFile
	}
	result, err := util.Command(commend)
	if err != nil {
		log.Fatalf("[bingo] ERROR test fail, please check your testPath !!, err: %v", err)
	}
	return result, nil
}
