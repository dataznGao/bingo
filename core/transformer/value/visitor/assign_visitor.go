package visitor

import (
	"github.com/dataznGao/go_drill/core/config"
	"github.com/dataznGao/go_drill/core/ds"
	"github.com/dataznGao/go_drill/core/transformer"
	"github.com/dataznGao/go_drill/util"
	"go/ast"
)

type ValueAssignVisitor struct {
	lp    *config.LocationPatternP
	value interface{}
	File  *ds.File
}

func (v *ValueAssignVisitor) Visit(node ast.Node) ast.Visitor {
	if stmt, ok := node.(*ast.AssignStmt); ok {
		for i, lh := range stmt.Lhs {
			switch lh.(type) {
			case *ast.SelectorExpr:
				se := lh.(*ast.SelectorExpr)
				if transformer.VariableCanInjure(v.lp, se.Sel.Name) {
					deepNode := *se
					if rh, ok := stmt.Rhs[i].(*ast.BasicLit); ok {
						rh.Value = util.StrVal(v.value)
					}
					if transformer.HasRunError(v.File) {
						se = &deepNode
						transformer.CreateFile(v.File)
					}
				}
			case *ast.Ident:
				se := lh.(*ast.Ident)
				if transformer.VariableCanInjure(v.lp, se.Name) {
					deepNode := *se
					if rh, ok := stmt.Rhs[i].(*ast.BasicLit); ok {
						rh.Value = util.StrVal(v.value)
					}
					if transformer.HasRunError(v.File) {
						se = &deepNode
						transformer.CreateFile(v.File)
					}
				}
			}
		}
	}
	return nil
}
