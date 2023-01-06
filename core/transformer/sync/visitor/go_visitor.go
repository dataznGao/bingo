package visitor

import (
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/core/ds"
	"github.com/dataznGao/bingo/core/transformer"
	"github.com/dataznGao/bingo/util"
	"go/ast"
	"go/token"
	"log"
)

type SyncGoVisitor struct {
	lp   *config.LocationPatternP
	call *ast.CallExpr
	File *ds.File
}

func (v *SyncGoVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.GoStmt:
		if transformer.VariablesCanInjure(v.lp, []string{"*"}) {
			stmt := node.(*ast.GoStmt)
			log.Printf("[bingo] INFO 变异位置: %v", util.GetNodeCode(stmt))
			deepNode := *stmt
			var token token.Pos
			stmt.Go = token
			v.call = stmt.Call
			if transformer.HasRunError(v.File) {
				stmt = &deepNode
				transformer.CreateFile(v.File)
				log.Printf("[bingo] INFO 变异位置: %v, 本次变异失败", util.GetNodeCode(stmt))
			} else {
				log.Printf("[bingo] INFO 成功变异为: %v", util.GetNodeCode(stmt))
			}
		}
	}
	return v
}
