package visitor

import (
	"github.com/dataznGao/go_drill/core/config"
	"github.com/dataznGao/go_drill/core/ds"
	"github.com/dataznGao/go_drill/core/transformer"
	"go/ast"
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
		if transformer.VariableCanInjure(v.lp, vari) {
			deepNode := *ss
			deleteBranch := -1
			for i, stmt := range ss.Body.List {
				if cc, ok := stmt.(*ast.CaseClause); ok {
					if cc.List == nil {
						deleteBranch = i
					}
				}
			}
			if deleteBranch != -1 {
				ss.Body.List = append(ss.Body.List[0:deleteBranch], ss.Body.List[deleteBranch+1:]...)
			}
			if transformer.HasRunError(v.File) {
				ss = &deepNode
				transformer.CreateFile(v.File)
			}
		}
	}

	return nil
}
