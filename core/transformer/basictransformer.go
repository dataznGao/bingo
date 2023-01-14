package transformer

import (
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/core/ds"
	"github.com/dataznGao/bingo/util"
	"go/ast"
	"strings"
)

type Transformer interface{ ToInjure() }

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
	return (planStruct == "*" || planStruct == "" || planStruct == structName) && (planFunc == "*" || planFunc == "" || planFunc == funcName) && util.CanPerform(lp.StructP.ActivationRate) && util.CanPerform(lp.MethodP.ActivationRate)
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
func HasRunError(file *ds.File) (string, bool) {
	if file.IsInjured && file.InputPath != file.OutputPath {
		file.OriInputPath = file.InputPath
		file.InputPath = file.OutputPath
	}
	err := CreateFile(file)
	if err != nil {
		panic(err)
	}
	newPath := ""
	if file.InputPath == file.OutputPath {
		newPath = util.CompareAndExchange(file.FileName, file.OutputPath, file.OriInputPath)
	} else {
		newPath = util.CompareAndExchange(file.FileName, file.OutputPath, file.InputPath)
	}
	command := "cd " + util.GetFather(newPath) + " && go build"
	_, err = util.Command(command)
	if err != nil {
		return newPath, true
	}
	file.IsInjured = true
	return newPath, false
}
func CreateFile(file *ds.File) error {
	code := util.GetFileCode(file.File)
	newPath := ""
	if file.InputPath == file.OutputPath {
		newPath = util.CompareAndExchange(file.FileName, file.OutputPath, file.OriInputPath)
	} else {
		newPath = util.CompareAndExchange(file.FileName, file.OutputPath, file.InputPath)
	}
	err := util.CreateFile(newPath, code)
	return err
}
