package visitor

import (
	"github.com/dataznGao/go_drill/core/config"
	"github.com/dataznGao/go_drill/core/ds"
	"github.com/dataznGao/go_drill/util"
	"go/ast"
)

type SyncObjectVisitor struct {
	Config          *config.FaultConfig
	locationPattern *config.LocationPatternP
	File            *ds.File
}

func (v *SyncObjectVisitor) Visit(node ast.Node) ast.Visitor {
	if f, ok := node.(*ast.File); ok {
		locatePackages := util.ShowLocatePackage(f.Name.Name, v.Config.LocationPatterns)
		if len(v.Config.LocationPatterns) == 0 || len(locatePackages) != 0 {
			for _, locatePackage := range locatePackages {
				can := util.CanPerform(locatePackage.PackageP.ActivationRate)
				if can {
					objs := f.Decls
					for _, object := range objs {
						if decl, ok := object.(*ast.FuncDecl); ok {
							funcVisitor := &SyncFuncVisitor{
								lp:   locatePackage,
								File: v.File,
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
