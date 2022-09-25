package visitor

import (
	"github.com/dataznGao/go_drill/config"
	"github.com/dataznGao/go_drill/transformer"
	"go/ast"
)

type ExceptionUnhandledCondVisitor struct {
	lp  *config.LocationPatternP
	can bool
}

func (v *ExceptionUnhandledCondVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.BinaryExpr:
		if transformer.VariablesCanInjure(v.lp, transformer.GetVariable(node.(*ast.BinaryExpr))) {
			if x, ok := node.(*ast.BinaryExpr).X.(*ast.Ident); ok {
				if x.Name == "err" {
					if y, ok := node.(*ast.BinaryExpr).Y.(*ast.Ident); ok {
						if y.Name == "nil" {
							v.can = true
						}
					}
				}
			}
		}
	}
	return v
}
