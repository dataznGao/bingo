package exception_unhandled

import (
	"github.com/dataznGao/go_drill/config"
	"github.com/dataznGao/go_drill/transformer/exception_unhandled/visitor"
	"go/ast"
)

type ExceptionUnhandledTransformer struct {
	File   *ast.File
	Config *config.FaultConfig
}

func (t *ExceptionUnhandledTransformer) ToInjure() {
	objVisitor := &visitor.ExceptionUnhandledObjectVisitor{Config: t.Config}
	ast.Walk(objVisitor, t.File)
}
