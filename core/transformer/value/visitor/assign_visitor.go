package visitor

import (
	"fmt"
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/core/ds"
	"github.com/dataznGao/bingo/core/transformer"
	"github.com/dataznGao/bingo/util"
	"go/ast"
	"log"
)

type ValueAssignVisitor struct {
	lp    *config.LocationPatternP
	value interface{}
	File  *ds.File
}

func (v *ValueAssignVisitor) Visit(node ast.Node) ast.Visitor {
	if stmt, ok := node.(*ast.AssignStmt); ok {
		for i, lh := range stmt.Lhs {
			switch lh.(type) {
			case *ast.SelectorExpr:
				se := lh.(*ast.SelectorExpr)
				if transformer.VariableCanInjure(v.lp, se.Sel.Name) {
					if rh, ok := stmt.Rhs[i].(*ast.BasicLit); ok {
						lo := fmt.Sprintf("[bingo] INFO 变异位置: %v\n%v\n", v.File.FileName, util.GetNodeCode(se))
						value := rh.Value
						rh.Value = util.StrVal(v.value)
						if newPath, has := transformer.HasRunError(v.File); has {
							rh.Value = value
							transformer.CreateFile(v.File)
						} else {
							log.Printf(lo)
							log.Printf("[bingo] INFO 变异位置: %v\n变异为: \n%v\n", newPath, util.GetNodeCode(se))
						}
					}
				}
			case *ast.Ident:
				se := lh.(*ast.Ident)
				if transformer.VariableCanInjure(v.lp, se.Name) {
					if rh, ok := stmt.Rhs[i].(*ast.BasicLit); ok {
						lo := fmt.Sprintf("[bingo] INFO 变异位置: %v\n%v\n", v.File.FileName, util.GetNodeCode(se))
						value := rh.Value
						rh.Value = util.StrVal(v.value)
						if newPath, has := transformer.HasRunError(v.File); has {
							rh.Value = value
							transformer.CreateFile(v.File)
						} else {
							log.Printf(lo)
							log.Printf("[bingo] INFO 变异位置: %v\n变异为: \n%v\n", newPath, util.GetNodeCode(se))
						}
					}
				}
			}
		}
	}
	return nil
}
