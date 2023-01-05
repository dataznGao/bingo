package exception_uncaught

import (
	"github.com/dataznGao/go_drill/core/config"
	"github.com/dataznGao/go_drill/core/ds"
	"github.com/dataznGao/go_drill/core/transformer/exception_uncaught/visitor"
	"go/ast"
)

type ExceptionUncaughtTransformer struct {
	File   *ds.File
	Config *config.FaultConfig
}

func (t *ExceptionUncaughtTransformer) ToInjure() {
	objVisitor := &visitor.ExceptionUncaughtObjectVisitor{Config: t.Config, File: t.File}
	ast.Walk(objVisitor, t.File.File)
}
