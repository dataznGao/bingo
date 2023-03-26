package bingo

func main() {
	env := CreateMutationEnv("/Users/misery/GolandProjects/bingo/scene", "/Users/misery/GolandProjects/bingo/scene1", "")
	// env.ConditionInversedFault("util(1/5).myStruct(1/3).myFunc(1/2).myVariable | mmap(3/4).*.*.a(3/4)")
	//env.NullFault("scene.*.*.a")
	//env.ValueFault("scene.*.*.a", "\"str\"")
	env.ExceptionUncaughtFault("*.*.*.*")
	//env.ExceptionUnhandledFault("scene.*.*.*")
	//env.SwitchMissDefaultFault("scene.*.*.a")
	//env.ValueFault("scene.*.*.a", "\"str\"")
	//env.AttributeReversoFault("scene.*.*.c", 10)

	env.SyncFault("rpc_scene.*.*.*")
	f := MutationPerformer{}
	err := f.SetEnv(env).Test()
	if err != nil {
		return
	}
}
