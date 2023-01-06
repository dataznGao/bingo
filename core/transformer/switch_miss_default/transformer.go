package switch_miss_default

import (
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/core/ds"
	"github.com/dataznGao/bingo/core/transformer/switch_miss_default/visitor"
	"go/ast"
)

type SwitchMissDefaultTransformer struct {
	File   *ds.File
	Config *config.FaultConfig
}

func (t *SwitchMissDefaultTransformer) ToInjure() {
	objVisitor := &visitor.ValueObjectVisitor{Config: t.Config, File: t.File}
	ast.Walk(objVisitor, t.File.File)
}
