package visitor

import (
	"github.com/dataznGao/go_drill/config"
	"go/ast"
)

type ReversoForVisitor struct {
	lp    *config.LocationPatternP
	value interface{}
}

func (v *ReversoForVisitor) Visit(node ast.Node) ast.Visitor {
	if fs, ok := node.(*ast.ForStmt); ok {
		for _, stmt := range fs.Body.List {
			switch stmt.(type) {
			case *ast.AssignStmt:
				stmt := stmt.(*ast.AssignStmt)
				visitor := &ReversoAssignVisitor{
					lp:    v.lp,
					value: v.value,
				}
				ast.Walk(visitor, stmt)
			}
		}
	} else if fs, ok := node.(*ast.AssignStmt); ok {
		visitor := &ReversoAssignVisitor{
			lp:    v.lp,
			value: v.value,
		}
		ast.Walk(visitor, fs)
	}
	return v

}
