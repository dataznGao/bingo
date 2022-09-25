package value

import (
	"github.com/dataznGao/go_drill/config"
	"github.com/dataznGao/go_drill/transformer/value/visitor"
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
