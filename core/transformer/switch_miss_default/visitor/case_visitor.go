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

type SwitchMissDefaultCaseVisitor struct {
	lp   *config.LocationPatternP
	File *ds.File
}

func (v *SwitchMissDefaultCaseVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.SwitchStmt:
		ss := node.(*ast.SwitchStmt)
		var vari string
		switch ss.Tag.(type) {
		case *ast.Ident:
			vari = ss.Tag.(*ast.Ident).Name
		case *ast.SelectorExpr:
			vari = ss.Tag.(*ast.SelectorExpr).Sel.Name
		}
		if transformer.VariableCanInjure(v.File, v.lp, vari) {
			deleteBranch := -1
			lo := fmt.Sprintf("[bingo] INFO 变异位置: %v\n%v\n", v.File.FileName, util.GetNodeCode(ss))
			for i, stmt := range ss.Body.List {
				if cc, ok := stmt.(*ast.CaseClause); ok {
					if cc.List == nil {
						deleteBranch = i
					}
				}
			}
			if deleteBranch != -1 {
				replica := util.CopyStmtList(ss.Body.List)
				if deleteBranch == len(ss.Body.List)-1 {
					ss.Body.List = ss.Body.List[0:deleteBranch]
				} else {
					ss.Body.List = append(ss.Body.List[0:deleteBranch], ss.Body.List[deleteBranch+1:]...)
				}
				if newPath, has := transformer.HasRunError(v.File); has {
					ss.Body.List = replica
					transformer.CreateFile(v.File)
				} else {
					log.Printf(lo)
					log.Printf("[bingo] INFO 变异位置: %v\n变异为: \n%v\n", newPath, util.GetNodeCode(ss))
				}
			}
		}
	}

	return nil
}
