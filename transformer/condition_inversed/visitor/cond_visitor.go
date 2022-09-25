package visitor

import (
	"fundrill_code_fault/config"
	"fundrill_code_fault/transformer"
	"go/ast"
	"go/token"
)

type ConditionInversedCondVisitor struct {
	lp *config.LocationPatternP
}

func (v *ConditionInversedCondVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.BinaryExpr:
		if transformer.VariablesCanInjure(v.lp, transformer.GetVariable(node.(*ast.BinaryExpr))) {
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
		}
	}
	return v
}
