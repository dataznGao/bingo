package visitor

import (
	"fundrill_code_fault/config"
	"go/ast"
)

type ValueCaseVisitor struct {
	lp    *config.LocationPatternP
	value interface{}
}

func (v *ValueCaseVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.AssignStmt:
		stmt := node.(*ast.AssignStmt)
		visitor := &ValueAssignVisitor{
			lp:    v.lp,
			value: v.value,
		}
		ast.Walk(visitor, stmt)
	}

	return v
}
