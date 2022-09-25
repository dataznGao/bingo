package visitor

import (
	"fundrill_code_fault/config"
	"fundrill_code_fault/transformer"
	"go/ast"
)

type ExceptionUncaughtFuncVisitor struct {
	lp *config.LocationPatternP
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
						lp: v.lp,
					}
					ast.Walk(visitor, ifStmt)
					if visitor.can {
						decl.Body.List = append(decl.Body.List[0:i], decl.Body.List[i+1:]...)
					}
				} else if caseStmt, ok := stmt.(*ast.SwitchStmt); ok {
					visitor := &ExceptionUncaughtIfVisitor{
						lp: v.lp,
					}
					ast.Walk(visitor, caseStmt)
				} else if asignStmt, ok := stmt.(*ast.AssignStmt); ok {
					visitor := &ExceptionUncaughtAssignVisitor{
						lp: v.lp,
					}
					ast.Walk(visitor, asignStmt)
				}
			}
		}
	}
	return nil
}
