package visitor

import (
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/core/ds"
	"github.com/dataznGao/bingo/core/transformer"
	"go/ast"
	"go/token"
)

type SyncGoVisitor struct {
	lp   *config.LocationPatternP
	call *ast.CallExpr
	File *ds.File
}

func (v *SyncGoVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.GoStmt:
		if transformer.VariablesCanInjure(v.lp, []string{"*"}) {
			stmt := node.(*ast.GoStmt)
			var token token.Pos
			stmt.Go = token
			v.call = stmt.Call
		}
	}
	return v
}
