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
			deepNode := *stmt
			log.Printf("[bingo] INFO 变异位置: %v", util.GetNodeCode(stmt))
			stmt.Body.List = make([]ast.Stmt, 0)
			if transformer.HasRunError(v.File) {
				stmt = &deepNode
				transformer.CreateFile(v.File)
			} else {
				log.Printf("[bingo] INFO 成功变异为: %v", util.GetNodeCode(stmt))
			}
		}
	}
	return nil
}
