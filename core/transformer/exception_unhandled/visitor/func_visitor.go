package visitor

import (
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/core/ds"
	"github.com/dataznGao/bingo/core/transformer"
	"go/ast"
)

type ExceptionUnhandledFuncVisitor struct {
	lp   *config.LocationPatternP
	File *ds.File
}

func (v *ExceptionUnhandledFuncVisitor) Visit(node ast.Node) ast.Visitor {
	if decl, ok := node.(*ast.FuncDecl); ok {
		var structs []*ast.Field
		if decl.Recv == nil || decl.Recv.List == nil {
			structs = make([]*ast.Field, 0)
		} else {
			structs = decl.Recv.List
		}
		can := transformer.FunCanInjure(v.lp, structs, decl.Name.Name)
		if can {
			for _, stmt := range decl.Body.List {
				if ifStmt, ok := stmt.(*ast.IfStmt); ok {
					visitor := &ExceptionUnhandledIfVisitor{lp: v.lp, can: false, File: v.File}
					ast.Walk(visitor, ifStmt)
				} else if caseStmt, ok := stmt.(*ast.SwitchStmt); ok {
					visitor := &ExceptionUnhandledIfVisitor{lp: v.lp, File: v.File}
					ast.Walk(visitor, caseStmt)
				}
			}
		}
	}
	return v
}
