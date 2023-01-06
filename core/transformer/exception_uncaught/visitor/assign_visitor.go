package visitor

import (
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/core/ds"
	"github.com/dataznGao/bingo/core/transformer"
	"github.com/dataznGao/bingo/util"
	"go/ast"
	"log"
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
					log.Printf("[bingo] INFO 变异位置: %v", util.GetNodeCode(lh))
					deepNode := *lh.(*ast.Ident)
					se.Name = "_"
					if transformer.HasRunError(v.File) {
						se = &deepNode
						transformer.CreateFile(v.File)
						log.Printf("[bingo] INFO 变异位置: %v, 本次变异失败", util.GetNodeCode(lh))
					} else {
						log.Printf("[bingo] INFO 成功变异为: %v", util.GetNodeCode(lh))
					}
				}

			}
		}
	}
	return nil
}
