package switch_miss_default

import (
	"github.com/dataznGao/go_drill/config"
	"github.com/dataznGao/go_drill/transformer/switch_miss_default/visitor"
	"go/ast"
)

type SwitchMissDefaultTransformer struct {
	File   *ast.File
	Config *config.FaultConfig
}

func (t *SwitchMissDefaultTransformer) ToInjure() {
	objVisitor := &visitor.ValueObjectVisitor{Config: t.Config}
	ast.Walk(objVisitor, t.File)
}
