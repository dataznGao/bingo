package condition_inversed

import (
	"fundrill_code_fault/config"
	"fundrill_code_fault/transformer/condition_inversed/visitor"
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
