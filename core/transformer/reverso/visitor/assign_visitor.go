package visitor

import (
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/core/ds"
	"github.com/dataznGao/bingo/core/transformer"
	"github.com/dataznGao/bingo/util"
	"go/ast"
	"log"
)

type ReversoAssignVisitor struct {
	lp    *config.LocationPatternP
	value interface{}
	File  *ds.File
}

func (v *ReversoAssignVisitor) Visit(node ast.Node) ast.Visitor {
	if stmt, ok := node.(*ast.AssignStmt); ok {

		for i, lh := range stmt.Lhs {
			switch lh.(type) {
			case *ast.SelectorExpr:
				se := lh.(*ast.SelectorExpr)
				if transformer.VariableCanInjure(v.lp, se.Sel.Name) {
					log.Printf("[bingo] INFO 变异位置: %v", util.GetNodeCode(lh))
					replicaNode := *stmt
					if rh, ok := stmt.Rhs[i].(*ast.BasicLit); ok {
						rh.Value = util.StrVal(v.value) + "*" + rh.Value
					}
					if transformer.HasRunError(v.File) {
						stmt = &replicaNode
						transformer.CreateFile(v.File)
						log.Printf("[bingo] INFO 变异位置: %v, 本次变异失败", util.GetNodeCode(lh))
					} else {
						log.Printf("[bingo] INFO 成功变异为: %v", util.GetNodeCode(lh))
					}
				}
			case *ast.Ident:
				se := lh.(*ast.Ident)
				if transformer.VariableCanInjure(v.lp, se.Name) {
					replicaNode := *stmt
					if rh, ok := stmt.Rhs[i].(*ast.BasicLit); ok {
						rh.Value = util.StrVal(v.value) + "*" + rh.Value
					}
					if transformer.HasRunError(v.File) {
						stmt = &replicaNode
						transformer.CreateFile(v.File)
					}
				}
			}
		}
	}
	return nil
}
