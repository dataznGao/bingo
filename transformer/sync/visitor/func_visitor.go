package visitor

import (
	"fundrill_code_fault/config"
	"fundrill_code_fault/transformer"
	"go/ast"
)

type SyncFuncVisitor struct {
	lp *config.LocationPatternP
}

func (v *SyncFuncVisitor) Visit(node ast.Node) ast.Visitor {
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
					visitor := &SyncIfVisitor{
						lp:  v.lp,
						can: false,
					}
					ast.Walk(visitor, ifStmt)
				} else if goStmt, ok := stmt.(*ast.GoStmt); ok {
					visitor := &SyncGoVisitor{
						lp:   v.lp,
						call: nil,
					}
					ast.Walk(visitor, goStmt)
					if visitor.call != nil {
						expr := new(ast.ExprStmt)
						expr.X = visitor.call
						decl.Body.List[i] = expr
					}
				} else if forStmt, ok := stmt.(*ast.ForStmt); ok {
					visitor := &SyncForVisitor{
						lp: v.lp,
					}
					ast.Walk(visitor, forStmt)
				} else if caseStmt, ok := stmt.(*ast.SwitchStmt); ok {
					visitor := &SyncCaseVisitor{
						lp: v.lp,
					}
					ast.Walk(visitor, caseStmt)
				}
			}
		}
	}
	return v
}
