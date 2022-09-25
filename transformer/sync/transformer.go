package sync

import (
	"github.com/dataznGao/go_drill/config"
	"github.com/dataznGao/go_drill/transformer/sync/visitor"
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
