package visitor

import (
	"fundrill_code_fault/config"
	"fundrill_code_fault/util"
	"go/ast"
)

type ExceptionUncaughtAssignVisitor struct {
	lp    *config.LocationPatternP
	value interface{}
}

func (v *ExceptionUncaughtAssignVisitor) Visit(node ast.Node) ast.Visitor {
	if stmt, ok := node.(*ast.AssignStmt); ok {
		for _, lh := range stmt.Lhs {
			switch lh.(type) {
			case *ast.Ident:
				se := lh.(*ast.Ident)
				if util.CanPerform(v.lp.VariableP.ActivationRate) {
					se.Name = "_"
				}
			}
		}
	}
	return nil
}
