package main

import (
	"github.com/dataznGao/go_drill/code_drill"
	"github.com/dataznGao/go_drill/env"
)

func main() {
	env := env.CreateFaultEnv("/Users/misery/GolandProjects/code_fault/mmap", "/Users/misery/GolandProjects/code_fault/mmap1")
	//env.ConditionInversedFault("util(1/5).myStruct(1/3).myFunc(1/2).myVariable | mmap(3/4).*.*.a(3/4)")
	//	ConditionInversedFault("util.myStruct.myFunc.myVariable")
	//
	//env.NullFault("mmap.*.*.a")
	//env.ValueFault("mmap.*.*.a", "\"str\"")
	//env.ExceptionUncaughtFault("mmap.*.*.*")
	//env.ExceptionUnhandledFault("mmap.*.*.*")
	//env.SwitchMissDefaultFault("mmap.*.*.a")
	env.SyncFault("mmap.*.*.*")
	env.ValueFault("mmap.*.*.a", "\"str\"")
	env.AttributeReversoFault("mmap.*.*.c", 10)
	f := code_drill.FaultPerformerFactory{}
	err := f.SetEnv(env).Run()
	if err != nil {
		return
	}
}
