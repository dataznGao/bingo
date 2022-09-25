package visitor

import (
	"github.com/dataznGao/go_drill/config"
	"github.com/dataznGao/go_drill/transformer"
	"go/ast"
)

type SwitchMissDefaultFuncVisitor struct {
	lp *config.LocationPatternP
}

func (v *SwitchMissDefaultFuncVisitor) Visit(node ast.Node) ast.Visitor {
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
				if caseStmt, ok := stmt.(*ast.SwitchStmt); ok {
					visitor := &SwitchMissDefaultCaseVisitor{
						lp: v.lp,
					}
					ast.Walk(visitor, caseStmt)
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
