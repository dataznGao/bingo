package visitor

import (
	"github.com/dataznGao/go_drill/config"
	"go/ast"
)

type SyncCaseVisitor struct {
	lp *config.LocationPatternP
}

func (v *SyncCaseVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.IfStmt:
		stmt := node.(*ast.IfStmt)
		visitor := &SyncIfVisitor{
			lp: v.lp,
		}
		ast.Walk(visitor, stmt)
	}

	return v
}
