package visitor

import (
	"github.com/dataznGao/go_drill/config"
	"go/ast"
)

type ReversoIfVisitor struct {
	lp    *config.LocationPatternP
	value interface{}
}

func (v *ReversoIfVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.AssignStmt:
		stmt := node.(*ast.AssignStmt)
		visitor := &ReversoAssignVisitor{
			lp:    v.lp,
			value: v.value,
		}
		ast.Walk(visitor, stmt)
	}
	return v
}
