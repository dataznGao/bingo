package visitor

import (
	"fundrill_code_fault/config"
	"go/ast"
)

type SyncIfVisitor struct {
	lp  *config.LocationPatternP
	can bool
}

func (v *SyncIfVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.IfStmt:
		stmt := node.(*ast.IfStmt)
		for i, s := range stmt.Body.List {
			switch s.(type) {
			case *ast.GoStmt:
				stm := s.(*ast.GoStmt)
				goVisitor := &SyncGoVisitor{
					lp:   v.lp,
					call: nil,
				}
				ast.Walk(goVisitor, stm)
				if goVisitor.call != nil {
					var expr = new(ast.ExprStmt)
					expr.X = goVisitor.call
					stmt.Body.List[i] = expr
				}
			}
		}
	}
	return v
}
