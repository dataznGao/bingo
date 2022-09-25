package exception_uncaught

import (
	"github.com/dataznGao/go_drill/config"
	"github.com/dataznGao/go_drill/transformer/exception_uncaught/visitor"
	"go/ast"
)

type ExceptionUncaughtTransformer struct {
	File   *ast.File
	Config *config.FaultConfig
}

func (t *ExceptionUncaughtTransformer) ToInjure() {
	objVisitor := &visitor.ExceptionUncaughtObjectVisitor{Config: t.Config}
	ast.Walk(objVisitor, t.File)
}
