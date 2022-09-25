package sync

import (
	"fundrill_code_fault/config"
	"fundrill_code_fault/transformer/sync/visitor"
	"go/ast"
)

type SyncTransformer struct {
	File   *ast.File
	Config *config.FaultConfig
}

func (t *SyncTransformer) ToInjure() {
	objVisitor := &visitor.SyncObjectVisitor{Config: t.Config}
	ast.Walk(objVisitor, t.File)
}
