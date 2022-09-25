package switch_miss_default

import (
	"fundrill_code_fault/config"
	"fundrill_code_fault/transformer/switch_miss_default/visitor"
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
