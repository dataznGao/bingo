package visitor

import (
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/core/ds"
	"go/ast"
)

type ValueCaseVisitor struct {
	lp   *config.LocationPatternP
	File *ds.File
}

func (v *ValueCaseVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.IfStmt:
		stmt := node.(*ast.IfStmt)
		visitor := &ConditionInversedIfVisitor{lp: v.lp, File: v.File}
		ast.Walk(visitor, stmt)
	}
	return v
}
