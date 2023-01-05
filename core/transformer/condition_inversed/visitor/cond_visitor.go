package visitor

import (
	"github.com/dataznGao/go_drill/core/config"
	"github.com/dataznGao/go_drill/core/ds"
	"github.com/dataznGao/go_drill/core/transformer"
	"go/ast"
	"go/token"
)

type ConditionInversedCondVisitor struct {
	lp   *config.LocationPatternP
	File *ds.File
}

func (v *ConditionInversedCondVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.BinaryExpr:
		if transformer.VariablesCanInjure(v.lp, transformer.GetVariable(node.(*ast.BinaryExpr))) {
			deepNode := *node.(*ast.BinaryExpr)
			op := node.(*ast.BinaryExpr).Op
			switch op {
			case token.LAND:
				op = token.LOR
			case token.GTR:
				op = token.LEQ
			case token.LOR:
				op = token.LOR
			case token.LEQ:
				op = token.GTR
			case token.LSS:
				op = token.GEQ
			case token.GEQ:
				op = token.LSS
			case token.NEQ:
				op = token.EQL
			case token.EQL:
				op = token.NEQ
			}
			node.(*ast.BinaryExpr).Op = op
			if transformer.HasRunError(v.File) {
				node = &deepNode
				transformer.CreateFile(v.File)
			}
		}

	}
	return v
}