package visitor

import (
	"github.com/dataznGao/go_drill/config"
	"go/ast"
)

type ExceptionUncaughtVisitor struct {
	lp *config.LocationPatternP
}

func (v *ExceptionUncaughtVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.IfStmt:
		stmt := node.(*ast.IfStmt)
		visitor := &ExceptionUncaughtIfVisitor{
			lp: v.lp,
		}
		ast.Walk(visitor, stmt)
	}
	return v
}
