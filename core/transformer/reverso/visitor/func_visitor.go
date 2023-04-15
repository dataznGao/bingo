package visitor

import (
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/core/ds"
	"github.com/dataznGao/bingo/core/transformer"
	"go/ast"
)

type ReversoFuncVisitor struct {
	lp    *config.LocationPatternP
	value interface{}
	File  *ds.File
}

func (v *ReversoFuncVisitor) Visit(node ast.Node) ast.Visitor {
	if decl, ok := node.(*ast.FuncDecl); ok {
		var structs []*ast.Field
		if decl.Recv == nil || decl.Recv.List == nil {
			structs = make([]*ast.Field, 0)
		} else {
			structs = decl.Recv.List
		}
		can := transformer.FunCanInjure(v.File, v.lp, structs, decl.Name.Name)
		if can {
			// 对函数段中不同的stmt进行单独处理
			for _, stmt := range decl.Body.List {
				if ifStmt, ok := stmt.(*ast.IfStmt); ok {
					visitor := &ReversoIfVisitor{
						lp:    v.lp,
						value: v.value,
						File:  v.File,
					}
					ast.Walk(visitor, ifStmt)
				} else if forStmt, ok := stmt.(*ast.ForStmt); ok {
					visitor := &ReversoForVisitor{
						lp:    v.lp,
						value: v.value,
						File:  v.File,
					}
					ast.Walk(visitor, forStmt)
				} else if caseStmt, ok := stmt.(*ast.SwitchStmt); ok {
					visitor := &ReversoCaseVisitor{
						lp:    v.lp,
						value: v.value,
						File:  v.File,
					}
					ast.Walk(visitor, caseStmt)
				}
			}
		}
	}
	return v
}
