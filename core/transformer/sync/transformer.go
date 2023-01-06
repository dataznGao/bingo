package sync

import (
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/core/ds"
	"github.com/dataznGao/bingo/core/transformer/sync/visitor"
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
