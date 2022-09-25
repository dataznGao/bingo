package visitor

import (
	"fundrill_code_fault/config"
	"fundrill_code_fault/transformer"
	"fundrill_code_fault/util"
	"go/ast"
)

type ReversoAssignVisitor struct {
	lp    *config.LocationPatternP
	value interface{}
}

func (v *ReversoAssignVisitor) Visit(node ast.Node) ast.Visitor {
	if stmt, ok := node.(*ast.AssignStmt); ok {
		for i, lh := range stmt.Lhs {
			switch lh.(type) {
			case *ast.SelectorExpr:
				se := lh.(*ast.SelectorExpr)
				if transformer.VariableCanInjure(v.lp, se.Sel.Name) {
					if rh, ok := stmt.Rhs[i].(*ast.BasicLit); ok {
						rh.Value = util.StrVal(v.value) + "*" + rh.Value
					}
				}
			case *ast.Ident:
				se := lh.(*ast.Ident)
				if transformer.VariableCanInjure(v.lp, se.Name) {
					if rh, ok := stmt.Rhs[i].(*ast.BasicLit); ok {
						rh.Value = util.StrVal(v.value) + "*" + rh.Value
					}
				}
			}
		}
	}
	return nil
}
