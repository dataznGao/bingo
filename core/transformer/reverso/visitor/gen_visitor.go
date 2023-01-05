package visitor

import (
	"github.com/dataznGao/go_drill/core/config"
	"github.com/dataznGao/go_drill/core/ds"
	"github.com/dataznGao/go_drill/core/transformer"
	"github.com/dataznGao/go_drill/util"
	"go/ast"
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
							deepNode := *sp
							if ident, ok := sp.Values[i].(*ast.Ident); ok {
								ident.Name = util.StrVal(v.value) + " * " + ident.Name
							} else if ident, ok := sp.Values[i].(*ast.BasicLit); ok {
								ident.Value = util.StrVal(v.value) + " * " + ident.Value
							}
							if transformer.HasRunError(v.File) {
								sp = &deepNode
								transformer.CreateFile(v.File)
							}
						}
					}
				}
			}
		}
	}
	return nil
}
