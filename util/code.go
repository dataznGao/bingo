package util

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func GetIfCode(node *ast.IfStmt) (string, error) {
	var output []byte
	buffer := bytes.NewBuffer(output)
	err := format.Node(buffer, token.NewFileSet(), node)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func GetFuncCode(node *ast.FuncDecl) string {
	var output []byte
	buffer := bytes.NewBuffer(output)
	format.Node(buffer, token.NewFileSet(), node)
	return buffer.String()
}

func GetAssignCode(node *ast.AssignStmt) string {
	var output []byte
	buffer := bytes.NewBuffer(output)
	format.Node(buffer, token.NewFileSet(), node)
	return buffer.String()
}

func GetNodeCode(node ast.Node) string {
	var output []byte
	buffer := bytes.NewBuffer(output)
	format.Node(buffer, token.NewFileSet(), node)
	return buffer.String()
}

func GetBuildInfo(fileName string) []byte {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("获取file失败, filename: %v", fileName)
	}
	fileStr := string(file)
	resStr := ""
	reg1 := regexp.MustCompile("//\\s*go:build")
	indexs := reg1.FindAllIndex(file, -1)
	reg2 := regexp.MustCompile("//\\s*\\+build")
	indexs = append(indexs, reg2.FindAllIndex(file, -1)...)
	n := len(fileStr)
	for _, index := range indexs {
		start := index[0]
		end := start
		for i := start; i < n; i++ {
			if i == n-1 || fileStr[i] == '\n' {
				end = i + 1
				break
			}
		}
		resStr += fileStr[start:end]
	}
	return []byte(resStr)
}

func GetFileCode(node *ast.File) []byte {
	var output []byte
	buffer := bytes.NewBuffer(output)
	format.Node(buffer, token.NewFileSet(), node)
	return buffer.Bytes()
}

func Copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func CopyStmtList(stmt []ast.Stmt) []ast.Stmt {
	replica := make([]ast.Stmt, len(stmt))
	copy(replica, stmt)
	return replica
}
