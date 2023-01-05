package visitor

import (
	"github.com/dataznGao/go_drill/core/config"
	"github.com/dataznGao/go_drill/core/ds"
	"go/ast"
)

type ValueIfVisitor struct {
	lp    *config.LocationPatternP
	value interface{}
	File  *ds.File
}

func (v *ValueIfVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.AssignStmt:
		stmt := node.(*ast.AssignStmt)
		visitor := &ValueAssignVisitor{
			lp:    v.lp,
			value: v.value,
			File:  v.File,
		}
		ast.Walk(visitor, stmt)
	}
	return v
}
