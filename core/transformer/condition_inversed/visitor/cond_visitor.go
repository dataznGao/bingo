package visitor

import (
	"fmt"
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/core/ds"
	"github.com/dataznGao/bingo/core/transformer"
	"github.com/dataznGao/bingo/util"
	"go/ast"
	"go/token"
	"log"
)

type ConditionInversedCondVisitor struct {
	lp   *config.LocationPatternP
	File *ds.File
}

func (v *ConditionInversedCondVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.BinaryExpr:
		if transformer.VariablesCanInjure(v.lp, transformer.GetVariable(node.(*ast.BinaryExpr))) {
			lo := fmt.Sprintf("[bingo] INFO 变异位置: %v\n%v\n", v.File.FileName, util.GetNodeCode(node))
			op := node.(*ast.BinaryExpr).Op
			switch op {
			case token.LAND:
				op = token.LOR
			case token.GTR:
				op = token.LEQ
			case token.LOR:
				op = token.LOR
			case token.LEQ:
				op = token.GTR
			case token.LSS:
				op = token.GEQ
			case token.GEQ:
				op = token.LSS
			case token.NEQ:
				op = token.EQL
			case token.EQL:
				op = token.NEQ
			}
			rep := node.(*ast.BinaryExpr).Op
			node.(*ast.BinaryExpr).Op = op
			if newPath, has := transformer.HasRunError(v.File); has {
				node.(*ast.BinaryExpr).Op = rep
				transformer.CreateFile(v.File)
			} else {
				log.Printf(lo)
				log.Printf("[bingo] INFO 变异位置: %v\n变异为: \n%v\n", newPath, util.GetNodeCode(node))
			}
		}

	}
	return v
}
