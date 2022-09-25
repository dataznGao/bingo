package visitor

import (
	"fundrill_code_fault/config"
	"fundrill_code_fault/transformer"
	"go/ast"
)

type ExceptionUnhandledFuncVisitor struct {
	lp *config.LocationPatternP
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
			// 对函数段中不同的stmt进行单独处理
			for _, stmt := range decl.Body.List {

				if ifStmt, ok := stmt.(*ast.IfStmt); ok {
					visitor := &ExceptionUnhandledIfVisitor{
						lp:  v.lp,
						can: false,
					}
					ast.Walk(visitor, ifStmt)

				} else if caseStmt, ok := stmt.(*ast.SwitchStmt); ok {
					visitor := &ExceptionUnhandledIfVisitor{
						lp: v.lp,
					}
					ast.Walk(visitor, caseStmt)
				}
			}
		}
	}
	return v
}
