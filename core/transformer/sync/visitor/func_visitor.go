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

type SyncFuncVisitor struct {
	lp   *config.LocationPatternP
	File *ds.File
}

func (v *SyncFuncVisitor) Visit(node ast.Node) ast.Visitor {
	if decl, ok := node.(*ast.FuncDecl); ok {
		var structs []*ast.Field
		if decl.Recv == nil || decl.Recv.List == nil {
			structs = make([]*ast.Field, 0)
		} else {
			structs = decl.Recv.List
		}
		can := transformer.FunCanInjure(v.lp, structs, decl.Name.Name)
		if can {
			// 对函数段中不同的stmt进行单独处理
			for i, stmt := range decl.Body.List {
				if ifStmt, ok := stmt.(*ast.IfStmt); ok {
					visitor := &SyncIfVisitor{
						lp:   v.lp,
						can:  false,
						File: v.File,
					}
					ast.Walk(visitor, ifStmt)
				} else if goStmt, ok := stmt.(*ast.GoStmt); ok {
					visitor := &SyncGoVisitor{
						lp:   v.lp,
						call: nil,
						File: v.File,
					}
					ast.Walk(visitor, goStmt)
					lo := fmt.Sprintf("[bingo] INFO 变异位置: %v\n%v\n", v.File.FileName, util.GetNodeCode(stmt))
					var expr = new(ast.ExprStmt)
					expr.X = visitor.call
					rep := decl.Body.List[i]
					decl.Body.List[i] = expr
					if newPath, has := transformer.HasRunError(v.File); has {
						decl.Body.List[i] = rep
						transformer.CreateFile(v.File)
					} else {
						log.Printf(lo)
						log.Printf("[bingo] INFO 变异位置: %v\n变异为: \n%v\n", newPath, util.GetNodeCode(expr))
					}
				} else if forStmt, ok := stmt.(*ast.ForStmt); ok {
					visitor := &SyncForVisitor{
						lp:   v.lp,
						File: v.File,
					}
					ast.Walk(visitor, forStmt)
				} else if caseStmt, ok := stmt.(*ast.SwitchStmt); ok {
					visitor := &SyncCaseVisitor{
						lp:   v.lp,
						File: v.File,
					}
					ast.Walk(visitor, caseStmt)
				}
			}
		}
	}
	return v
}
