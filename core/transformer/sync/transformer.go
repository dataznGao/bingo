package sync

import (
	"github.com/dataznGao/go_drill/core/config"
	"github.com/dataznGao/go_drill/core/ds"
	"github.com/dataznGao/go_drill/core/transformer/sync/visitor"
	"go/ast"
)

type SyncTransformer struct {
	File   *ds.File
	Config *config.FaultConfig
}

func (t *SyncTransformer) ToInjure() {
	objVisitor := &visitor.SyncObjectVisitor{Config: t.Config, File: t.File}
	ast.Walk(objVisitor, t.File.File)
}
