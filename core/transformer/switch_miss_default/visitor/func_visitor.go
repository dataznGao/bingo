package visitor

import (
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/core/ds"
	"github.com/dataznGao/bingo/core/transformer"
	"go/ast"
)

type SwitchMissDefaultFuncVisitor struct {
	lp   *config.LocationPatternP
	File *ds.File
}

func (v *SwitchMissDefaultFuncVisitor) Visit(node ast.Node) ast.Visitor {
	if decl, ok := node.(*ast.FuncDecl); ok {
		var structs []*ast.Field
		if decl.Recv == nil || decl.Recv.List == nil {
			structs = make([]*ast.Field, 0)
		} else {
			structs = decl.Recv.List
		}
		can := transformer.FunCanInjure(v.File, v.lp, structs, decl.Name.Name)
		if can {
			if decl.Body != nil && decl.Body.List != nil {
				// 对函数段中不同的stmt进行单独处理
				for _, stmt := range decl.Body.List {
					if caseStmt, ok := stmt.(*ast.SwitchStmt); ok {
						visitor := &SwitchMissDefaultCaseVisitor{
							lp:   v.lp,
							File: v.File,
						}
						ast.Walk(visitor, caseStmt)
					}
				}
			}
		}
	} else if decl, ok := node.(*ast.SwitchStmt); ok {
		visitor := &SwitchMissDefaultCaseVisitor{
			lp: v.lp,
		}
		ast.Walk(visitor, decl)
	}
	return v
}
