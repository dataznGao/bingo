package value

import (
	"fundrill_code_fault/config"
	"fundrill_code_fault/transformer/value/visitor"
	"go/ast"
)

type ValueTransformer struct {
	File   *ast.File
	Config *config.FaultConfig
}

func (t *ValueTransformer) ToInjure() {
	objVisitor := &visitor.ValueObjectVisitor{Config: t.Config}
	ast.Walk(objVisitor, t.File)
}
