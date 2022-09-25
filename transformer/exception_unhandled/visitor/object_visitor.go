package visitor

import (
	"github.com/dataznGao/go_drill/config"
	"github.com/dataznGao/go_drill/util"
	"go/ast"
)

type ExceptionUnhandledObjectVisitor struct {
	Config          *config.FaultConfig
	locationPattern *config.LocationPatternP
}

func (v *ExceptionUnhandledObjectVisitor) Visit(node ast.Node) ast.Visitor {
	if f, ok := node.(*ast.File); ok {
		locatePackages := util.ShowLocatePackage(f.Name.Name, v.Config.LocationPatterns)
		if len(v.Config.LocationPatterns) == 0 || len(locatePackages) != 0 {
			for _, locatePackage := range locatePackages {
				can := util.CanPerform(locatePackage.PackageP.ActivationRate)
				if can {
					objs := f.Decls
					for _, object := range objs {
						if decl, ok := object.(*ast.FuncDecl); ok {
							funcVisitor := &ExceptionUnhandledFuncVisitor{
								lp: locatePackage,
							}
							ast.Walk(funcVisitor, decl)
						}
					}
				}
			}
		}
	}
	return nil
}
