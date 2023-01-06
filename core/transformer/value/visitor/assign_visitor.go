package visitor

import (
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
					log.Printf("[bingo] INFO 变异位置: %v", util.GetNodeCode(se))
					deepNode := *se
					if rh, ok := stmt.Rhs[i].(*ast.BasicLit); ok {
						rh.Value = util.StrVal(v.value)
					}
					if transformer.HasRunError(v.File) {
						se = &deepNode
						transformer.CreateFile(v.File)
						log.Printf("[bingo] INFO 变异位置: %v, 本次变异失败", util.GetNodeCode(se))
					} else {
						log.Printf("[bingo] INFO 成功变异为: %v", util.GetNodeCode(se))
					}
				}
			case *ast.Ident:
				se := lh.(*ast.Ident)
				if transformer.VariableCanInjure(v.lp, se.Name) {
					log.Printf("[bingo] INFO 变异位置: %v", util.GetNodeCode(se))
					deepNode := *se
					if rh, ok := stmt.Rhs[i].(*ast.BasicLit); ok {
						rh.Value = util.StrVal(v.value)
					}
					if transformer.HasRunError(v.File) {
						se = &deepNode
						transformer.CreateFile(v.File)
						log.Printf("[bingo] INFO 变异位置: %v, 本次变异失败", util.GetNodeCode(se))
					} else {
						log.Printf("[bingo] INFO 成功变异为: %v", util.GetNodeCode(se))
					}
				}
			}
		}
	}
	return nil
}
