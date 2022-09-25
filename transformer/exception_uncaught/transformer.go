package exception_uncaught

import (
	"fundrill_code_fault/config"
	"fundrill_code_fault/transformer/exception_uncaught/visitor"
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
