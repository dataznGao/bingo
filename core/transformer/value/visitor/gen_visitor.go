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

type ValueGenVisitor struct {
	lp    *config.LocationPatternP
	value interface{}
	File  *ds.File
}

func (v *ValueGenVisitor) Visit(node ast.Node) ast.Visitor {
	if gen, ok := node.(*ast.GenDecl); ok {
		for _, spec := range gen.Specs {
			switch spec.(type) {
			case *ast.ValueSpec:
				sp := spec.(*ast.ValueSpec)
				for i, name := range sp.Names {
					if name.Name == v.lp.VariableP.Name {
						can := util.CanPerform(v.lp.VariableP.ActivationRate)
						if can {
							lo := fmt.Sprintf("[bingo] INFO 变异位置: %v\n%v\n", v.File.FileName, util.GetNodeCode(gen))
							if sp.Values == nil {
								sp.Values = []ast.Expr{ast.NewIdent(util.StrVal(v.value))}
								if newPath, has := transformer.HasRunError(v.File); has {
									sp.Values = nil
									transformer.CreateFile(v.File)
								} else {
									log.Printf(lo)
									log.Printf("[bingo] INFO 变异位置: %v\n变异为: \n%v\n", newPath, util.GetNodeCode(gen))
								}
							} else if ident, ok := sp.Values[i].(*ast.Ident); ok {
								nam := ident.Name
								ident.Name = util.StrVal(v.value)
								if newPath, has := transformer.HasRunError(v.File); has {
									ident.Name = nam
									transformer.CreateFile(v.File)
								} else {
									log.Printf(lo)
									log.Printf("[bingo] INFO 变异位置: %v\n变异为: \n%v\n", newPath, util.GetNodeCode(gen))
								}
							}
						}
					}
				}
			}
		}
	}
	return nil
}
