package visitor

import (
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/core/ds"
	"github.com/dataznGao/bingo/core/transformer"
	"go/ast"
)

type ExceptionUncaughtFuncVisitor struct {
	lp   *config.LocationPatternP
	File *ds.File
}

func (v *ExceptionUncaughtFuncVisitor) Visit(node ast.Node) ast.Visitor {
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
			for i, stmt := range decl.Body.List {
				if ifStmt, ok := stmt.(*ast.IfStmt); ok {
					visitor := &ExceptionUncaughtIfVisitor{
						lp:   v.lp,
						File: v.File,
					}
					ast.Walk(visitor, ifStmt)
					if visitor.can {
						if i == len(decl.Body.List)-1 {
							decl.Body.List = decl.Body.List[0:i]
						} else {
							decl.Body.List = append(decl.Body.List[0:i], decl.Body.List[i+1:]...)
						}
					}
				} else if caseStmt, ok := stmt.(*ast.SwitchStmt); ok {
					visitor := &ExceptionUncaughtIfVisitor{
						lp:   v.lp,
						File: v.File,
					}
					ast.Walk(visitor, caseStmt)
				} else if asignStmt, ok := stmt.(*ast.AssignStmt); ok {
					visitor := &ExceptionUncaughtAssignVisitor{
						lp:   v.lp,
						File: v.File,
					}
					ast.Walk(visitor, asignStmt)
				}
			}
		}
	}
	return nil
}
