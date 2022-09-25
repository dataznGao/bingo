package visitor

import (
	"fundrill_code_fault/config"
	"fundrill_code_fault/util"
	"go/ast"
)

type ValueObjectVisitor struct {
	Config          *config.FaultConfig
	locationPattern *config.LocationPatternP
}

func (v *ValueObjectVisitor) Visit(node ast.Node) ast.Visitor {
	if f, ok := node.(*ast.File); ok {
		locatePackages := util.ShowLocatePackage(f.Name.Name, v.Config.LocationPatterns)
		if len(v.Config.LocationPatterns) == 0 || len(locatePackages) != 0 {
			for _, locatePackage := range locatePackages {
				can := util.CanPerform(locatePackage.PackageP.ActivationRate)
				if can {
					objs := f.Decls
					for _, object := range objs {
						if locatePackage.MethodP.Name == "" || locatePackage.MethodP.Name == "*" {
							if decl, ok := object.(*ast.GenDecl); ok {
								genVisitor := &ValueGenVisitor{
									lp:    locatePackage,
									value: v.Config.TargetValue,
								}
								ast.Walk(genVisitor, decl)
							}
						}
						if decl, ok := object.(*ast.FuncDecl); ok {
							funcVisitor := &ValueFuncVisitor{
								lp:    locatePackage,
								value: v.Config.TargetValue,
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
