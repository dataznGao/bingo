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
					if se.Name == "err" {
						log.Printf("[bingo] INFO 变异位置: %v\n%v\n", v.File.FileName, util.GetNodeCode(stmt))
						se.Name = "_"
						if newPath, has := transformer.HasRunError(v.File); has {
							se.Name = "err"
							transformer.CreateFile(v.File)
							log.Printf("[bingo] INFO 变异位置: %v\n%v\n本次变异失败\n", newPath, util.GetNodeCode(stmt))
						} else {
							log.Printf("[bingo] INFO 变异位置: %v\n成功变异为: \n%v\n", newPath, util.GetNodeCode(stmt))
						}
					}
				}

			}
		}
	}
	return nil
}
