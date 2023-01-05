package reverso

import (
	"github.com/dataznGao/go_drill/core/config"
	"github.com/dataznGao/go_drill/core/ds"
	"github.com/dataznGao/go_drill/core/transformer/reverso/visitor"
	"go/ast"
)

type AttributeReversoTransformer struct {
	File   *ds.File
	Config *config.FaultConfig
}

func (t *AttributeReversoTransformer) ToInjure() {
	objVisitor := &visitor.ReversoObjectVisitor{Config: t.Config, File: t.File}
	ast.Walk(objVisitor, t.File.File)
}
