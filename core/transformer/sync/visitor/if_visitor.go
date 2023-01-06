package visitor

import (
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/core/ds"
	"go/ast"
)

type SyncIfVisitor struct {
	lp   *config.LocationPatternP
	can  bool
	File *ds.File
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
					File: v.File,
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
