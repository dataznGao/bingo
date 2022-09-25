package null

import (
	"fundrill_code_fault/config"
	"fundrill_code_fault/transformer/value/visitor"
	"go/ast"
)

type NullTransformer struct {
	File   *ast.File
	Config *config.FaultConfig
}

func (t *NullTransformer) ToInjure() {
	objVisitor := &visitor.ValueObjectVisitor{Config: t.Config}
	ast.Walk(objVisitor, t.File)
}
