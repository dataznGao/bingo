package transformer

import (
	"github.com/dataznGao/bingo/constant"
	"github.com/dataznGao/bingo/core/clean"
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/core/ds"
	"github.com/dataznGao/bingo/util"
	"go/ast"
	"strings"
)

var ShouldPrint bool

var OutInfo [][]string

func init() {
	OutInfo = make([][]string, 0)
	OutInfo = append(OutInfo, []string{"file_name", "package_name", "struct_name", "function_name"})
}

type Transformer interface {
	ToInjure()
}

func FunCanInjure(file *ds.File, lp *config.LocationPatternP, structs []*ast.Field, funcName string) bool {
	if len(structs) == 0 {
		if oneCanInjure(file, lp, "", funcName) {
			return true
		}
	}
	for _, struc := range structs {
		if x, ok := struc.Type.(*ast.StarExpr); ok {
			if ide, ok := x.X.(*ast.Ident); ok {
				if oneCanInjure(file, lp, ide.Name, funcName) {
					return true
				}
			}
		} else if ide, ok := struc.Type.(*ast.Ident); ok {
			if oneCanInjure(file, lp, ide.Name, funcName) {
				return true
			}
		}
	}
	return false
}

func oneCanInjure(file *ds.File, lp *config.LocationPatternP, structName, funcName string) bool {
	planStruct := strings.TrimSpace(lp.StructP.Name)
	planFunc := strings.TrimSpace(lp.MethodP.Name)
	can := (planStruct == "*" || planStruct == "" || planStruct == structName) &&
		(planFunc == "*" || planFunc == "" || planFunc == funcName) &&
		util.CanPerform(lp.StructP.ActivationRate) && util.CanPerform(lp.MethodP.ActivationRate)
	if can && ShouldPrint {
		if file.Info == nil {
			file.Info = &ds.PrintInfo{
				PackageName:  file.File.Name.Name,
				StructName:   "",
				FuncName:     "",
				VariableName: "",
				FileName:     file.FileName,
			}
		}
		file.Info.StructName = structName
		file.Info.FuncName = funcName
		file.Info.FileName = file.FileName
	}
	return can
}

func VariablesCanInjure(file *ds.File, lp *config.LocationPatternP, variables []string) bool {
	has := util.Contains(lp.VariableP.Name, variables)
	if has {
		can := util.CanPerform(lp.VariableP.ActivationRate)
		if can && ShouldPrint {
			if file.Info == nil {
				file.Info = &ds.PrintInfo{
					PackageName:  "",
					StructName:   "",
					FuncName:     "",
					VariableName: "",
				}
			}
			file.Info.VariableName = lp.VariableP.Name
			file.Info.FileName = file.FileName
		}
		return can
	} else {
		return false
	}
}

func VariableCanInjure(file *ds.File, lp *config.LocationPatternP, variable string) bool {
	has := lp.VariableP.Name == variable
	if has {
		can := util.CanPerform(lp.VariableP.ActivationRate)
		if can && ShouldPrint {
			if file.Info == nil {
				file.Info = &ds.PrintInfo{
					PackageName:  "",
					StructName:   "",
					FuncName:     "",
					VariableName: "",
					FileName:     file.FileName,
				}
			}
			file.Info.VariableName = lp.VariableP.Name
			file.Info.FileName = file.FileName
		}
		return can
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
	command := "cd " + file.OutputPath + " && go mod tidy"
	_, err = util.Command(command)
	if err != nil {
		return newPath, true
	}
	command = "cd " + util.GetFather(newPath) + " && go build"
	res, err := util.Command(command)
	hasErr := false
	if err != nil {
		cleaner := clean.ErrCleaner{Err: res, FileName: newPath}
		err = cleaner.Fix()
		// 修理了还是有错，给true
		if err != nil {
			hasErr = true
		}
	}
	if hasErr {
		constant.InjuredFailureCnt += 1
	} else {
		constant.InjuredSuccessCnt += 1
	}
	file.IsInjured = true
	return newPath, hasErr
}

func CreateFile(file *ds.File) error {
	// 将trace也打印出来
	code := util.GetFileCode(file.File)
	if OutInfo == nil {
		OutInfo = make([][]string, 0)
		OutInfo = append(OutInfo, []string{"file_name", "package_name", "struct_name", "function_name"})
	}
	OutInfo = append(OutInfo, []string{file.Info.FileName,
		file.Info.PackageName, file.Info.StructName, file.Info.FuncName})
	newPath := ""
	if file.InputPath == file.OutputPath {
		newPath = util.CompareAndExchange(file.FileName, file.OutputPath, file.OriInputPath)
	} else {
		newPath = util.CompareAndExchange(file.FileName, file.OutputPath, file.InputPath)
	}
	err := util.CreateFile(newPath, code)
	return err
}
