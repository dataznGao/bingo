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

type ExceptionUncaughtFuncVisitor struct {
	lp   *config.LocationPatternP
	File *ds.File
}

func (v *ExceptionUncaughtFuncVisitor) Visit(node ast.Node) ast.Visitor {
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
			n := len(decl.Body.List)
			for i := n - 1; i >= 0; i-- {
				stmt := decl.Body.List[i]
				if ifStmt, ok := stmt.(*ast.IfStmt); ok {
					visitor := &ExceptionUncaughtIfVisitor{
						lp:   v.lp,
						File: v.File,
					}
					ast.Walk(visitor, ifStmt)
					if visitor.can {
						lo := fmt.Sprintf("[bingo] INFO 变异位置: %v\n%v\n", v.File.FileName, util.GetNodeCode(decl))
						replica := util.CopyStmtList(decl.Body.List)
						if i == len(decl.Body.List)-1 {
							decl.Body.List = decl.Body.List[0:i]
						} else {
							decl.Body.List = append(decl.Body.List[0:i], decl.Body.List[i+1:]...)
						}
						if newPath, has := transformer.HasRunError(v.File); has {
							decl.Body.List = replica
							transformer.CreateFile(v.File)
						} else {
							log.Printf(lo)
							log.Printf("[bingo] INFO 变异位置: %v\n变异为: \n%v\n", newPath, util.GetNodeCode(decl))
						}
					}
				} else if caseStmt, ok := stmt.(*ast.SwitchStmt); ok {
					visitor := &ExceptionUncaughtIfVisitor{
						lp:   v.lp,
						File: v.File,
					}
					ast.Walk(visitor, caseStmt)
				} else if asignStmt, ok := stmt.(*ast.AssignStmt); ok {
					visitor := &ExceptionUncaughtAssignVisitor{
						lp:   v.lp,
						File: v.File,
					}
					ast.Walk(visitor, asignStmt)
				}
			}
		}
	}
	return nil
}
