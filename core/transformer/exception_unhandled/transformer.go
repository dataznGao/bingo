package exception_unhandled

import (
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/core/ds"
	"github.com/dataznGao/bingo/core/transformer/exception_unhandled/visitor"
	"go/ast"
)

type ExceptionUnhandledTransformer struct {
	File   *ds.File
	Config *config.FaultConfig
}

func (t *ExceptionUnhandledTransformer) ToInjure() {
	objVisitor := &visitor.ExceptionUnhandledObjectVisitor{Config: t.Config, File: t.File}
	ast.Walk(objVisitor, t.File.File)
}
