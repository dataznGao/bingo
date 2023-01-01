package main

func main() {
	env := CreateFaultEnv("/Users/misery/GolandProjects/code_fault/scene", "/Users/misery/GolandProjects/code_fault/scene1")
	// env.ConditionUnversedFault("util(1/5).myStruct(1/3).myFunc(1/2).myVariable | mmap(3/4).*.*.a(3/4)")
	//env.NullFault("scene.*.*.a")
	env.ValueFault("scene.*.*.a", "\"str\"")
	//env.ExceptionUncaughtFault("scene.*.*.*")
	//env.ExceptionUnhandledFault("scene.*.*.*")
	//env.SwitchMissDefaultFault("scene.*.*.a")
	//env.ValueFault("scene.*.*.a", "\"str\"")
	//env.AttributeReversoFault("scene.*.*.c", 10)

	env.SyncFault("rpc_scene.*.*.*")
	f := FaultPerformerFactory{}
	err := f.SetEnv(env).Run()
	if err != nil {
		return
	}
}
