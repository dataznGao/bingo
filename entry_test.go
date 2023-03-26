package bingo

import "testing"

func TestFaultPerformerFactory_Run(t *testing.T) {
	env := CreateMutationEnv("/Users/misery/GolandProjects/jupiter",
		"/Users/misery/GolandProjects/errr",
		"/Users/misery/GolandProjects/jupiter/test/e2e")
	// env.ConditionInversedFault("util(1/5).myStruct(1/3).myFunc(1/2).myVariable | mmap(3/4).*.*.a(3/4)")
	//env.NullFault("*.*.*.a")
	//env.ConditionInversedFault("*.*.*.*")
	//env.ExceptionUncaughtFault("*.*.*.*")
	//env.ExceptionUncaughtFault("*.*.*.*")
	env.ExceptionUnhandledFault("*.*.*.*")
	env.SwitchMissDefaultFault("scene.*.*.a")
	env.ValueFault("*.*.*.id", "\"str\"")
	env.AttributeReversoFault("*.*.*.c", 10)

	//env.SyncFault("*.*.*.*")
	f := MutationPerformer{}
	err := f.SetEnv(env).Test()
	f.SetEnv(env)
	if err != nil {
		return
	}
}
