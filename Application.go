package bingo

func main() {
	env := CreateMutationEnv("/Users/misery/GolandProjects/bingo/scene", "/Users/misery/GolandProjects/bingo/scene1", "")
	env.ExceptionUncaughtFault("*.*.*.*")
	env.SyncFault("rpc_scene.*.*.*")
	f := MutationPerformer{}
	err := f.SetEnv(env).Run()
	if err != nil {
		return
	}
}
