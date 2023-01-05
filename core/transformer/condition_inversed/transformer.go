package condition_inversed

import (
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/core/ds"
	"github.com/dataznGao/bingo/core/transformer/condition_inversed/visitor"
	"go/ast"
)

type ConditionInversedTransformer struct {
	File   *ds.File
	Config *config.FaultConfig
}

func (t *ConditionInversedTransformer) ToInjure() {
	objVisitor := &visitor.ConditionInversedObjectVisitor{Config: t.Config, File: t.File}
	ast.Walk(objVisitor, t.File.File)
}
