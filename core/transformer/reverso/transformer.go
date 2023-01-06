package reverso

import (
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/core/ds"
	"github.com/dataznGao/bingo/core/transformer/reverso/visitor"
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
