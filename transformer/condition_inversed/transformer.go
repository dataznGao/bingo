package condition_inversed

import (
	"github.com/dataznGao/go_drill/config"
	"github.com/dataznGao/go_drill/transformer/condition_inversed/visitor"
	"go/ast"
)

type ConditionInversedTransformer struct {
	File   *ast.File
	Config *config.FaultConfig
}

func (t *ConditionInversedTransformer) ToInjure() {
	objVisitor := &visitor.ConditionInversedObjectVisitor{Config: t.Config}
	ast.Walk(objVisitor, t.File)
}
