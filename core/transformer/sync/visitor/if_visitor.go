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

type SyncIfVisitor struct {
	lp   *config.LocationPatternP
	can  bool
	File *ds.File
}

func (v *SyncIfVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.IfStmt:
		stmt := node.(*ast.IfStmt)
		for i, s := range stmt.Body.List {
			switch s.(type) {
			case *ast.GoStmt:
				stm := s.(*ast.GoStmt)
				goVisitor := &SyncGoVisitor{lp: v.lp, call: nil, File: v.File}
				ast.Walk(goVisitor, stm)
				lo := fmt.Sprintf("[bingo] INFO 变异位置: %v\n%v\n", v.File.FileName, util.GetNodeCode(s))
				var expr = new(ast.ExprStmt)
				expr.X = goVisitor.call
				rep := stmt.Body.List[i]
				stmt.Body.List[i] = expr
				if newPath, has := transformer.HasRunError(v.File); has {
					stmt.Body.List[i] = rep
					transformer.CreateFile(v.File)
				} else {
					log.Printf(lo)
					log.Printf("[bingo] INFO 变异位置: %v\n变异为: \n%v\n", newPath, util.GetNodeCode(expr))
				}
			}
		}
	}
	return v
}
