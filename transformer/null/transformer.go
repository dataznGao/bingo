package null

import (
	"github.com/dataznGao/go_drill/config"
	"github.com/dataznGao/go_drill/transformer/value/visitor"
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
