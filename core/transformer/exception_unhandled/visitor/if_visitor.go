package visitor

import (
	"github.com/dataznGao/go_drill/core/config"
	"github.com/dataznGao/go_drill/core/ds"
	"github.com/dataznGao/go_drill/core/transformer"
	"go/ast"
)

type ExceptionUnhandledIfVisitor struct {
	lp   *config.LocationPatternP
	can  bool
	File *ds.File
}

func (v *ExceptionUnhandledIfVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.IfStmt:
		stmt := node.(*ast.IfStmt)
		deepNode := *stmt
		condVisitor := &ExceptionUnhandledCondVisitor{
			lp:   v.lp,
			can:  false,
			File: v.File,
		}
		ast.Walk(condVisitor, stmt)
		if condVisitor.can {
			stmt.Body.List = make([]ast.Stmt, 0)
		}
		if transformer.HasRunError(v.File) {
			stmt = &deepNode
			transformer.CreateFile(v.File)
		}
	}
	return nil
}
