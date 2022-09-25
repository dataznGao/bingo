package ds

import "go/ast"

type FileInjure struct {
	File      *ast.File
	IsInjured bool
}
