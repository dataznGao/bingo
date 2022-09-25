package transformer

import (
	"fundrill_code_fault/config"
	"fundrill_code_fault/util"
	"go/ast"
	"strings"
)

type Transformer interface {
	ToInjure()
}

func FunCanInjure(lp *config.LocationPatternP, structs []*ast.Field, funcName string) bool {
	if len(structs) == 0 {
		if oneCanInjure(lp, "", funcName) {
			return true
		}
	}
	for _, struc := range structs {
		if x, ok := struc.Type.(*ast.StarExpr); ok {
			if ide, ok := x.X.(*ast.Ident); ok {
				if oneCanInjure(lp, ide.Name, funcName) {
					return true
				}
			}
		} else if ide, ok := struc.Type.(*ast.Ident); ok {
			if oneCanInjure(lp, ide.Name, funcName) {
				return true
			}
		}
	}
	return false
}

func oneCanInjure(lp *config.LocationPatternP, structName, funcName string) bool {
	planStruct := strings.TrimSpace(lp.StructP.Name)
	planFunc := strings.TrimSpace(lp.MethodP.Name)
	return (planStruct == "*" || planStruct == "" || planStruct == structName) &&
		(planFunc == "*" || planFunc == "" || planFunc == funcName) &&
		util.CanPerform(lp.StructP.ActivationRate) && util.CanPerform(lp.MethodP.ActivationRate)
}

func VariablesCanInjure(lp *config.LocationPatternP, variables []string) bool {
	has := util.Contains(lp.VariableP.Name, variables)
	if has {
		return util.CanPerform(lp.VariableP.ActivationRate)
	} else {
		return false
	}
}

func VariableCanInjure(lp *config.LocationPatternP, variable string) bool {
	has := lp.VariableP.Name == variable
	if has {
		return util.CanPerform(lp.VariableP.ActivationRate)
	} else {
		return false
	}
}

func GetVariable(node *ast.BinaryExpr) []string {
	res := make([]string, 0)
	q := &util.Queue[*ast.BinaryExpr]{}
	q.Offer(node)
	for !q.IsEmpty() {
		poll := q.Poll()
		if bin, ok := poll.X.(*ast.BinaryExpr); ok {
			q.Offer(bin)
		} else if bin, ok := poll.X.(*ast.Ident); ok {
			res = append(res, bin.Name)
		} else if se, ok := poll.X.(*ast.SelectorExpr); ok {
			res = append(res, se.Sel.Name)
		}
		if bin, ok := poll.Y.(*ast.BinaryExpr); ok {
			q.Offer(bin)
		} else if bin, ok := poll.Y.(*ast.Ident); ok {
			res = append(res, bin.Name)
		} else if se, ok := poll.Y.(*ast.SelectorExpr); ok {
			res = append(res, se.Sel.Name)
		}
	}
	return res
}
