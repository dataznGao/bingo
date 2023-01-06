package visitor

import (
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/core/ds"
	"github.com/dataznGao/bingo/core/transformer"
	"go/ast"
)

type ValueFuncVisitor struct {
	lp    *config.LocationPatternP
	value interface{}
	File  *ds.File
}

func (v *ValueFuncVisitor) Visit(node ast.Node) ast.Visitor {
	if decl, ok := node.(*ast.FuncDecl); ok {
		var structs []*ast.Field
		if decl.Recv == nil || decl.Recv.List == nil {
			structs = make([]*ast.Field, 0)
		} else {
			structs = decl.Recv.List
		}
		can := transformer.FunCanInjure(v.lp, structs, decl.Name.Name)
		if can {
			// 对函数段中不同的stmt进行单独处理
			for _, stmt := range decl.Body.List {
				if ifStmt, ok := stmt.(*ast.IfStmt); ok {
					visitor := &ValueIfVisitor{
						lp:    v.lp,
						value: v.value,
						File:  v.File,
					}
					ast.Walk(visitor, ifStmt)
				} else if forStmt, ok := stmt.(*ast.ForStmt); ok {
					visitor := &ValueForVisitor{
						lp:    v.lp,
						value: v.value,
						File:  v.File,
					}
					ast.Walk(visitor, forStmt)
				} else if caseStmt, ok := stmt.(*ast.SwitchStmt); ok {
					visitor := &ValueCaseVisitor{
						lp:    v.lp,
						value: v.value,
						File:  v.File,
					}
					ast.Walk(visitor, caseStmt)
				} else if assignStmt, ok := stmt.(*ast.AssignStmt); ok {
					visitor := &ValueAssignVisitor{
						lp:    v.lp,
						value: v.value,
						File:  v.File,
					}
					ast.Walk(visitor, assignStmt)
				}
			}
		}
	}
	return v
}
