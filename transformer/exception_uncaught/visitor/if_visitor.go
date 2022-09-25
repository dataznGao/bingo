package visitor

import (
	"github.com/dataznGao/go_drill/config"
	"go/ast"
)

type ExceptionUncaughtIfVisitor struct {
	lp  *config.LocationPatternP
	can bool
}

func (v *ExceptionUncaughtIfVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.IfStmt:
		stmt := node.(*ast.IfStmt)
		condVisitor := &ExceptionUncaughtCondVisitor{
			lp:  v.lp,
			can: false,
		}
		ast.Walk(condVisitor, stmt)
		if condVisitor.can {
			v.can = true
		}
	case *ast.AssignStmt:
		stmt := node.(*ast.AssignStmt)
		visitor := &ExceptionUncaughtAssignVisitor{
			lp: v.lp,
		}
		ast.Walk(visitor, stmt)
	}
	return v
}
