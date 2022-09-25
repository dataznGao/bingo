package visitor

import (
	"fundrill_code_fault/config"
	"go/ast"
)

type ExceptionUnhandledIfVisitor struct {
	lp  *config.LocationPatternP
	can bool
}

func (v *ExceptionUnhandledIfVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.IfStmt:
		stmt := node.(*ast.IfStmt)
		condVisitor := &ExceptionUnhandledCondVisitor{
			lp:  v.lp,
			can: false,
		}
		ast.Walk(condVisitor, stmt)
		if condVisitor.can {
			stmt.Body.List = make([]ast.Stmt, 0)
		}
	}
	return nil
}
