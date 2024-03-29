package bingo

import (
	"testing"
)

func TestFaultEnv_ValueFault(t *testing.T) {
	var l LocationPattern = "util(1/5).myStruct(1/3).myFunc(1/2).myVariable | main.      .*.*"
	l.parse()
}

func TestCreateFaultEnv(t *testing.T) {
	env := CreateMutationEnv("/Users/misery/GolandProjects/bingo/mmap",
		"/Users/misery/GolandProjects/bingo/mmap1", "")
	env.ValueFault("util(1/5).myStruct(1/3).myFunc(1/2).myVariable | main.*.*.*", nil).
		ValueFault("util(1/5).myStruct(1/3).myFunc(1/2).myVariable", 1)

}
