package bingo

import "testing"

func TestFaultPerformerFactory_Run(t *testing.T) {
	env := CreateMutationEnv("/Users/misery/GolandProjects/jupiter",
		"/Users/misery/GolandProjects/tmp_enhance",
		"/Users/misery/GolandProjects/jupiter/test/e2e")
	//env.ConditionInversedFault("*.*.*.*")
	//env.NullFault("*.*.*.a")
	//env.ConditionInversedFault("*.*.*.*")
	env.ExceptionUncaughtFault("*.*.*.*")
	env.ExceptionUncaughtFault("*.*.*.*")
	env.ExceptionUnhandledFault("*.*.*.*")
	env.SwitchMissDefaultFault("scene.*.*.a")
	env.ValueFault("*.*.*.id", "\"str\"")
	env.AttributeReversoFault("*.*.*.c", 10)

	env.SyncFault("*.*.*.*")
	f := MutationPerformer{}
	res, err := f.SetEnv(env).Run(true)
	println(res)
	if err != nil {
		return
	}
}
