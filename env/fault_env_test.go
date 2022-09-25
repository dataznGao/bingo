package env

import "testing"

func TestFaultEnv_ValueFault(t *testing.T) {
	var l LocationPattern = "util(1/5).myStruct(1/3).myFunc(1/2).myVariable | main.      .*.*"
	l.parse()
}

func TestCreateFaultEnv(t *testing.T) {
	env := CreateFaultEnv("/Users/misery/GolandProjects/code_fault/mmap", "/Users/misery/GolandProjects/code_fault/mmap1")
	env.ValueFault("util(1/5).myStruct(1/3).myFunc(1/2).myVariable | main.      .*.*", nil).
		ValueFault("util(1/5).myStruct(1/3).myFunc(1/2).myVariable", 1)

}
