package visitor

import (
	"github.com/dataznGao/go_drill/config"
	"go/ast"
)

type SyncForVisitor struct {
	lp    *config.LocationPatternP
	value interface{}
}

func (v *SyncForVisitor) Visit(node ast.Node) ast.Visitor {
	if fs, ok := node.(*ast.ForStmt); ok {
		for i, stmt := range fs.Body.List {
			switch stmt.(type) {
			case *ast.GoStmt:
				stm := stmt.(*ast.GoStmt)
				goVisitor := &SyncGoVisitor{
					lp:   v.lp,
					call: nil,
				}
				ast.Walk(goVisitor, stm)
				if goVisitor.call != nil {
					var expr = new(ast.ExprStmt)
					expr.X = goVisitor.call
					fs.Body.List[i] = expr
				}
			}
		}
	}
	return v

}
