package visitor

import (
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/core/ds"
	"go/ast"
)

type ValueForVisitor struct {
	lp    *config.LocationPatternP
	value interface{}
	File  *ds.File
}

func (v *ValueForVisitor) Visit(node ast.Node) ast.Visitor {
	if fs, ok := node.(*ast.ForStmt); ok {
		for _, stmt := range fs.Body.List {
			switch stmt.(type) {
			case *ast.AssignStmt:
				stmt := stmt.(*ast.AssignStmt)
				visitor := &ValueAssignVisitor{
					lp:    v.lp,
					value: v.value,
					File:  v.File,
				}
				ast.Walk(visitor, stmt)
			}
		}
	} else if fs, ok := node.(*ast.AssignStmt); ok {
		visitor := &ValueAssignVisitor{
			lp:    v.lp,
			value: v.value,
			File:  v.File,
		}
		ast.Walk(visitor, fs)
	}
	return v

}
