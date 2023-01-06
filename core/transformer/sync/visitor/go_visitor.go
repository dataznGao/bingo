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
			log.Printf("[bingo] INFO 变异位置: %v\n%v\n", v.File.FileName, util.GetNodeCode(stmt))
			var token token.Pos
			stmt.Go = token
			v.call = stmt.Call
			if newPath, has := transformer.HasRunError(v.File); has {
				v.call = nil
				transformer.CreateFile(v.File)
				log.Printf("[bingo] INFO 变异位置: %v\n%v\n本次变异失败\n", newPath, util.GetNodeCode(stmt))
			} else {
				log.Printf("[bingo] INFO 变异位置: %v\n成功变异为: \n%v\n", newPath, util.GetNodeCode(v.call))
			}
		}
	}
	return v
}
