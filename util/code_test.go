package util

import "testing"

func TestGetBuildInfo(t *testing.T) {
	GetBuildInfo("/Users/misery/GolandProjects/bingo/scene/rpc_scene/rpc_test.go")
}

func TestInsertFileHead(t *testing.T) {
	err := InsertFileHead("/Users/misery/GolandProjects/bingo/scene/rpc_scene/rpc_test.go", []byte("11111"))
	println(err)
}
