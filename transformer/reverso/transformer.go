package reverso

import (
	"fundrill_code_fault/config"
	"fundrill_code_fault/transformer/reverso/visitor"
	"go/ast"
)

type AttributeReversoTransformer struct {
	File   *ast.File
	Config *config.FaultConfig
}

func (t *AttributeReversoTransformer) ToInjure() {
	objVisitor := &visitor.ReversoObjectVisitor{Config: t.Config}
	ast.Walk(objVisitor, t.File)
}
