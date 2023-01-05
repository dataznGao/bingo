package visitor

import (
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/core/ds"
	"go/ast"
)

type ExceptionUncaughtIfVisitor struct {
	lp   *config.LocationPatternP
	can  bool
	File *ds.File
}

func (v *ExceptionUncaughtIfVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.IfStmt:
		stmt := node.(*ast.IfStmt)
		condVisitor := &ExceptionUncaughtCondVisitor{
			lp:   v.lp,
			can:  false,
			File: v.File,
		}
		ast.Walk(condVisitor, stmt)
		if condVisitor.can {
			v.can = true
		}
	case *ast.AssignStmt:
		stmt := node.(*ast.AssignStmt)
		visitor := &ExceptionUncaughtAssignVisitor{
			lp:   v.lp,
			File: v.File,
		}
		ast.Walk(visitor, stmt)
	}
	return v
}
