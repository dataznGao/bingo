package visitor

import (
	"github.com/dataznGao/go_drill/core/config"
	"github.com/dataznGao/go_drill/core/ds"
	"github.com/dataznGao/go_drill/core/transformer"
	"github.com/dataznGao/go_drill/util"
	"go/ast"
)

type ReversoAssignVisitor struct {
	lp    *config.LocationPatternP
	value interface{}
	File  *ds.File
}

func (v *ReversoAssignVisitor) Visit(node ast.Node) ast.Visitor {
	if stmt, ok := node.(*ast.AssignStmt); ok {

		for i, lh := range stmt.Lhs {
			switch lh.(type) {
			case *ast.SelectorExpr:
				se := lh.(*ast.SelectorExpr)
				if transformer.VariableCanInjure(v.lp, se.Sel.Name) {
					replicaNode := *stmt
					if rh, ok := stmt.Rhs[i].(*ast.BasicLit); ok {
						rh.Value = util.StrVal(v.value) + "*" + rh.Value
					}
					if transformer.HasRunError(v.File) {
						stmt = &replicaNode
						transformer.CreateFile(v.File)
					}
				}
			case *ast.Ident:
				se := lh.(*ast.Ident)
				if transformer.VariableCanInjure(v.lp, se.Name) {
					replicaNode := *stmt
					if rh, ok := stmt.Rhs[i].(*ast.BasicLit); ok {
						rh.Value = util.StrVal(v.value) + "*" + rh.Value
					}
					if transformer.HasRunError(v.File) {
						stmt = &replicaNode
						transformer.CreateFile(v.File)
					}
				}
			}
		}
	}
	return nil
}
