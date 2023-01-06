package visitor

import (
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/core/ds"
	"github.com/dataznGao/bingo/core/transformer"
	"go/ast"
)

type ExceptionUncaughtCondVisitor struct {
	can  bool
	lp   *config.LocationPatternP
	File *ds.File
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
