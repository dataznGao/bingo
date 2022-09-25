package exception_unhandled

import (
	"fundrill_code_fault/config"
	"fundrill_code_fault/transformer/exception_unhandled/visitor"
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
