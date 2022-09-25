package visitor

import (
	"fundrill_code_fault/config"
	"fundrill_code_fault/util"
	"go/ast"
)

type ValueGenVisitor struct {
	lp    *config.LocationPatternP
	value interface{}
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
							if ident, ok := sp.Values[i].(*ast.Ident); ok {
								ident.Name = util.StrVal(v.value)
							}
						}
					}
				}
			}
		}
	}
	return nil
}
