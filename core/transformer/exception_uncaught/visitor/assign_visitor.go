package visitor

import (
	"github.com/dataznGao/go_drill/core/config"
	"github.com/dataznGao/go_drill/core/ds"
	"github.com/dataznGao/go_drill/core/transformer"
	"github.com/dataznGao/go_drill/util"
	"go/ast"
)

type ExceptionUncaughtAssignVisitor struct {
	lp    *config.LocationPatternP
	value interface{}
	File  *ds.File
}

func (v *ExceptionUncaughtAssignVisitor) Visit(node ast.Node) ast.Visitor {
	if stmt, ok := node.(*ast.AssignStmt); ok {
		for _, lh := range stmt.Lhs {
			switch lh.(type) {
			case *ast.Ident:
				se := lh.(*ast.Ident)
				if util.CanPerform(v.lp.VariableP.ActivationRate) {
					deepNode := *lh.(*ast.Ident)
					se.Name = "_"
					if transformer.HasRunError(v.File) {
						se = &deepNode
						transformer.CreateFile(v.File)
					}
				}

			}
		}
	}
	return nil
}
