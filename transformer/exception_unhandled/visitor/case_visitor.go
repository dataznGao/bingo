package visitor

import (
	"github.com/dataznGao/go_drill/config"
	"go/ast"
)

type ExceptionUnhandledVisitor struct {
	lp *config.LocationPatternP
}

func (v *ExceptionUnhandledVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.IfStmt:
		stmt := node.(*ast.IfStmt)
		visitor := &ExceptionUnhandledIfVisitor{
			lp: v.lp,
		}
		ast.Walk(visitor, stmt)
	}

	return v
}
