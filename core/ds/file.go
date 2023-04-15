package ds

import (
	"go/ast"
)

type FileInjure struct {
	File      *ast.File
	IsInjured bool
}

type File struct {
	File         *ast.File
	FileName     string
	InputPath    string
	OutputPath   string
	IsInjured    bool
	OriInputPath string
	Info         *PrintInfo
}

type PrintInfo struct {
	FileName     string
	PackageName  string
	StructName   string
	FuncName     string
	VariableName string
}
