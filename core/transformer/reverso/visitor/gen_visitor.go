package visitor

import (
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/core/ds"
	"github.com/dataznGao/bingo/core/transformer"
	"github.com/dataznGao/bingo/util"
	"go/ast"
	"log"
)

type ReversoGenVisitor struct {
	lp    *config.LocationPatternP
	value interface{}
	File  *ds.File
}

func (v *ReversoGenVisitor) Visit(node ast.Node) ast.Visitor {
	if gen, ok := node.(*ast.GenDecl); ok {
		for _, spec := range gen.Specs {
			switch spec.(type) {
			case *ast.ValueSpec:
				sp := spec.(*ast.ValueSpec)
				for i, name := range sp.Names {
					if name.Name == v.lp.VariableP.Name {
						can := util.CanPerform(v.lp.VariableP.ActivationRate)
						if can {
							log.Printf("[bingo] INFO 变异位置: %v\n%v\n", v.File.FileName, util.GetNodeCode(sp))
							if ident, ok := sp.Values[i].(*ast.Ident); ok {
								name := ident.Name
								ident.Name = util.StrVal(v.value) + " * " + ident.Name
								if newPath, has := transformer.HasRunError(v.File); has {
									ident.Name = name
									transformer.CreateFile(v.File)
									log.Printf("[bingo] INFO 变异位置: %v\n%v\n本次变异失败\n", newPath, util.GetNodeCode(sp))
								} else {
									log.Printf("[bingo] INFO 变异位置: %v\n成功变异为: \n%v\n", newPath, util.GetNodeCode(sp))
								}
							} else if ident, ok := sp.Values[i].(*ast.BasicLit); ok {
								value := ident.Value
								ident.Value = util.StrVal(v.value) + " * " + ident.Value
								if newPath, has := transformer.HasRunError(v.File); has {
									ident.Value = value
									transformer.CreateFile(v.File)
									log.Printf("[bingo] INFO 变异位置: %v\n%v \n本次变异失败\n", newPath, util.GetNodeCode(sp))
								} else {
									log.Printf("[bingo] INFO 变异位置: %v\n成功变异为: \n%v\n", newPath, util.GetNodeCode(sp))
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
