package reverso

import (
	"github.com/dataznGao/go_drill/config"
	"github.com/dataznGao/go_drill/transformer/reverso/visitor"
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
