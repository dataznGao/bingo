package visitor

import (
	"fundrill_code_fault/config"
	"go/ast"
)

type ValueCaseVisitor struct {
	lp *config.LocationPatternP
}

func (v *ValueCaseVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.IfStmt:
		stmt := node.(*ast.IfStmt)
		visitor := &ConditionInversedIfVisitor{
			lp: v.lp,
		}
		ast.Walk(visitor, stmt)
	}

	return v
}
