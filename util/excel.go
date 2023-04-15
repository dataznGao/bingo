package util

import (
	"bufio"
	"bytes"
	"github.com/tealeg/xlsx"
	"os"
)

func DataToExcel(fileName string, data [][]string) error {
	length := len(fileName)
	pos := length
	for pos = length - 1; pos >= 0; pos-- {
		if fileName[pos] == '/' {
			break
		}
	}
	if fileName == "" {
		return nil
	}
	prefix := fileName[:pos]
	if !isExist(prefix) {
		err := os.MkdirAll(prefix, os.ModePerm)
		if err != nil {
			return err
		}
	}
	excelFile := xlsx.NewFile()
	sheet, err := excelFile.AddSheet("1")
	if err != nil {
		return err
	}
	for _, datum := range data {
		row := sheet.AddRow()
		for _, elem := range datum {
			row.AddCell().SetString(elem)
		}
	}
	buf := new(bytes.Buffer)
	excelFile.Write(buf)
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer file.Close()
	writer := bufio.NewWriter(file)
	if err != nil {
		return err
	}
	_, err = writer.Write(buf.Bytes())
	return writer.Flush()
}
