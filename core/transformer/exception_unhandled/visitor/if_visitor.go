package visitor

import (
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/core/ds"
	"github.com/dataznGao/bingo/core/transformer"
	"github.com/dataznGao/bingo/util"
	"go/ast"
	"log"
)

type ExceptionUnhandledIfVisitor struct {
	lp   *config.LocationPatternP
	can  bool
	File *ds.File
}

func (v *ExceptionUnhandledIfVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.IfStmt:
		stmt := node.(*ast.IfStmt)
		condVisitor := &ExceptionUnhandledCondVisitor{
			lp:   v.lp,
			can:  false,
			File: v.File,
		}
		ast.Walk(condVisitor, stmt)
		if condVisitor.can {
			log.Printf("[bingo] INFO 变异位置: %v\n%v\n", v.File.FileName, util.GetNodeCode(stmt))
			replica := util.CopyStmtList(stmt.Body.List)
			stmt.Body.List = make([]ast.Stmt, 0)
			if newPath, has := transformer.HasRunError(v.File); has {
				stmt.Body.List = replica
				transformer.CreateFile(v.File)
				log.Printf("[bingo] INFO 变异位置: %v\n%v\n本次变异失败\n", newPath, util.GetNodeCode(stmt))
			} else {
				log.Printf("[bingo] INFO 变异位置: %v\n成功变异为: \n%v\n", newPath, util.GetNodeCode(stmt))
			}
		}
	}
	return nil
}
