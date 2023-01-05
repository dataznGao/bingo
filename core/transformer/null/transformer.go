package null

import (
	"github.com/dataznGao/go_drill/core/config"
	"github.com/dataznGao/go_drill/core/ds"
	"github.com/dataznGao/go_drill/core/transformer/value/visitor"
	"go/ast"
)

type NullTransformer struct {
	File   *ds.File
	Config *config.FaultConfig
}

func (t *NullTransformer) ToInjure() {
	objVisitor := &visitor.ValueObjectVisitor{Config: t.Config, File: t.File}
	ast.Walk(objVisitor, t.File.File)
}
