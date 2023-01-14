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

type SyncForVisitor struct {
	lp    *config.LocationPatternP
	value interface{}
	File  *ds.File
}

func (v *SyncForVisitor) Visit(node ast.Node) ast.Visitor {
	if fs, ok := node.(*ast.ForStmt); ok {
		for i, stmt := range fs.Body.List {
			switch stmt.(type) {
			case *ast.GoStmt:
				stm := stmt.(*ast.GoStmt)
				goVisitor := &SyncGoVisitor{lp: v.lp, call: nil, File: v.File}
				ast.Walk(goVisitor, stm)
				if goVisitor.call != nil {
					lo := fmt.Sprintf("[bingo] INFO 变异位置: %v\n%v\n", v.File.FileName, util.GetNodeCode(stmt))
					var expr = new(ast.ExprStmt)
					expr.X = goVisitor.call
					rep := fs.Body.List[i]
					fs.Body.List[i] = expr
					if newPath, has := transformer.HasRunError(v.File); has {
						fs.Body.List[i] = rep
						transformer.CreateFile(v.File)
					} else {
						log.Printf(lo)
						log.Printf("[bingo] INFO 变异位置: %v\n变异为: \n%v\n", newPath, util.GetNodeCode(expr))
					}
				}
			}
		}
	}
	return v
}
