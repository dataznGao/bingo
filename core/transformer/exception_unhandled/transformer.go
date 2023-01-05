package exception_unhandled

import (
	"github.com/dataznGao/go_drill/core/config"
	"github.com/dataznGao/go_drill/core/ds"
	"github.com/dataznGao/go_drill/core/transformer/exception_unhandled/visitor"
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
