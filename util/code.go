package util

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"io"
	"os"
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
