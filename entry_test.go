package bingo

import "testing"

func TestFaultPerformerFactory_Run(t *testing.T) {
	env := CreateFaultEnv("/Users/misery/GolandProjects/bingo/scene", "/Users/misery/GolandProjects/bingo/scene1")
	// env.ConditionInversedFault("util(1/5).myStruct(1/3).myFunc(1/2).myVariable | mmap(3/4).*.*.a(3/4)")
	//env.NullFault("scene.*.*.a")
	//env.ValueFault("dal.*.*.Db", "\"str\"")
	env.ConditionInversedFault("*.*.*.*")
	env.ExceptionUncaughtFault("*.*.*.*")
	//env.ExceptionUncaughtFault("scene.*.*.*")
	//env.ExceptionUnhandledFault("scene.*.*.*")
	//env.SwitchMissDefaultFault("scene.*.*.a")
	//env.ValueFault("*.*.*.id", "\"str\"")
	//env.AttributeReversoFault("scene.*.*.c", 10)

	//env.SyncFault("rpc_scene.*.*.*")
	f := FaultPerformerFactory{}
	err := f.SetEnv(env).Run()
	if err != nil {
		return
	}
}
