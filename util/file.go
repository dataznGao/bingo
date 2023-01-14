package util

import (
	"bufio"
	"go/ast"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func LoadAllFile(path string) ([]string, []string, error) {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, nil, err
	}
	res := make([]string, 0)
	notGoFile := make([]string, 0)
	err = getAllFile(path, dir, &res, &notGoFile)
	if err != nil {
		return nil, nil, err
	}
	return res, notGoFile, nil
}
func getAllFile(parent string, dir []fs.FileInfo, res, notGoFile *[]string) error {
	for _, file := range dir {
		absolutePath := path.Join(parent, file.Name())
		if file.IsDir() {
			readDir, err := ioutil.ReadDir(absolutePath)
			if err != nil {
				return err
			}
			err = getAllFile(absolutePath, readDir, res, notGoFile)
			if err != nil {
				return err
			}
		} else {
			if strings.HasSuffix(file.Name(), ".go") && !strings.HasSuffix(file.Name(), "test.go") {
				*res = append(*res, absolutePath)
			} else {
				*notGoFile = append(*notGoFile, absolutePath)
			}
		}
	}
	return nil
}
func ConvertConfigMap(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = ConvertConfigMap(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = ConvertConfigMap(v)
		}
	}
	return i
}
func GetFilePackage(file *ast.File) string {
	return file.Name.Name
}
func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}
func CreateFile(path string, code []byte) error {
	length := len(path)
	pos := length
	for pos = length - 1; pos >= 0; pos-- {
		if path[pos] == '/' {
			break
		}
	}
	if path == "" {
		return nil
	}
	prefix := path[:pos]
	if !isExist(prefix) {
		err := os.MkdirAll(prefix, os.ModePerm)
		if err != nil {
			return err
		}
	}
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer file.Close()
	writer := bufio.NewWriter(file)
	if err != nil {
		return err
	}
	_, err = writer.Write(code)
	writer.Flush()
	return err
}
func Clean(gc []string) error {
	for _, path := range gc {
		os.Remove(path)
	}
	return nil
}
func InsertFileHead(fileName string, info []byte) error {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	var buf = make([]byte, 0)
	buf = append(buf, info...)
	buf = append(buf, '\n')
	buf = append(buf, file...)
	return CreateFile(fileName, buf)
}
func GetFather(fileName string) string {
	n := len(fileName)
	end := n
	for i := n - 1; i >= 0; i-- {
		if fileName[i] == '/' {
			end = i
			break
		}
	}
	return fileName[:end]
}
