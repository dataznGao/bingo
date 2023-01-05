package visitor

import (
	"github.com/dataznGao/go_drill/core/config"
	"github.com/dataznGao/go_drill/core/ds"
	"github.com/dataznGao/go_drill/core/transformer"
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
			deepNode := *stmt
			var token token.Pos
			stmt.Go = token
			v.call = stmt.Call
			if transformer.HasRunError(v.File) {
				stmt = &deepNode
				transformer.CreateFile(v.File)
			}
		}
	}
	return v
}
