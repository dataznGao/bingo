package visitor

import (
	"fundrill_code_fault/config"
	"fundrill_code_fault/transformer"
	"go/ast"
)

type ExceptionUncaughtCondVisitor struct {
	can bool
	lp  *config.LocationPatternP
}

func (v *ExceptionUncaughtCondVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.BinaryExpr:
		if transformer.VariablesCanInjure(v.lp, transformer.GetVariable(node.(*ast.BinaryExpr))) {
			if x, ok := node.(*ast.BinaryExpr).X.(*ast.Ident); ok {
				if x.Name == "err" {
					v.can = true
				}
			}
		}
	}
	return v
}
