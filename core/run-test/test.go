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
	testFiles, err := util.LoadAllTestFile(testPath)
	if err != nil {
		return "", err
	}
	res := ""
	for _, file := range testFiles {
		res += "========================================================================\n" +
			"test  " + file + "\n" + "test result:\n" + "\n\n"
		one, err2 := testOne(file, inputPath)
		if err2 != nil {
			log.Fatalf("[bingo] ERROR test %v err, please check your testPath", file)
			return "", err2
		}
		res += one + "\n\n\n\n"
	}
	return res, nil
}

func testOne(testFile string, inputPath string) (string, error) {
	// cd /Users/misery/GolandProjects/bingo && go test -v -cover scene/double_write_scene/entry_test.go scene/double_write_scene/entry.go
	// 校验testPath
	if !strings.HasSuffix(testFile, "_test.go") {
		return "", errors.New("the testFile is not end with _test.go")
	}
	if !strings.HasPrefix(testFile, inputPath) {
		return "", errors.New("the testFile or inputPath set err! please check! err")
	}
	// 获取相对路径
	relativeTestPath := util.CompareAndExchange(testFile, "", inputPath+constant.Separator)
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
	result, err := util.CommandTest(commend)
	if err != nil {
		log.Fatalf("[bingo] ERROR test fail, please check your testFile !!, err: %v", err)
	}
	return result, nil
}
