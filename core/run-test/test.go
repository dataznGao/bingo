package run

import (
	"errors"
	"github.com/dataznGao/bingo/util"
	"log"
	"strings"
)

func Test(testPath string, inputPath string) (string, error) {
	// cd /Users/misery/GolandProjects/bingo && go test -v -cover scene/double_write_scene/entry_test.go scene/double_write_scene/entry.go
	// 校验testPath
	if !strings.HasPrefix(testPath, inputPath) {
		return "", errors.New("the testFile or inputPath set err! please check! err")
	}
	// 先把依赖搞定
	commend := "cd " + inputPath + " && go mod tidy"
	result, err := util.Command(commend)
	if err != nil {
		log.Fatalf("[bingo] ERROR go mod tidy fail, please check your inputPath !!, err: %v", err)
	}
	commend = "cd " + testPath + " && go test -v -cover"
	result, err = util.CommandTest(commend)
	if err != nil {
		log.Fatalf("[bingo] ERROR test fail, please check your testFile !!, err: %v", err)
	}
	result += "\n\n\n"
	return result, nil
}
