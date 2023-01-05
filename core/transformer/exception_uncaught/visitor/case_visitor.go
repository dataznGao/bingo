package visitor

import (
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/core/ds"
	"go/ast"
)

type ExceptionUncaughtVisitor struct {
	lp   *config.LocationPatternP
	File *ds.File
}

func (v *ExceptionUncaughtVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.IfStmt:
		stmt := node.(*ast.IfStmt)
		visitor := &ExceptionUncaughtIfVisitor{
			lp:   v.lp,
			File: v.File,
		}
		ast.Walk(visitor, stmt)
	}
	return v
}
