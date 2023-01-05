package visitor

import (
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/core/ds"
	"go/ast"
)

type ExceptionUnhandledVisitor struct {
	lp   *config.LocationPatternP
	File *ds.File
}

func (v *ExceptionUnhandledVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.IfStmt:
		stmt := node.(*ast.IfStmt)
		visitor := &ExceptionUnhandledIfVisitor{
			lp:   v.lp,
			File: v.File,
		}
		ast.Walk(visitor, stmt)
	}

	return v
}
