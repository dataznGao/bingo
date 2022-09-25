package code_drill

import (
	"fundrill_code_fault/env"
	"fundrill_code_fault/util"
	"math/rand"
	"testing"
)

func TestLoadPackage(t *testing.T) {
	files, _ := LoadPackage(Config.InputPath)
	for _, file := range files {
		for _, config := range Config.FaultPoints {
			transformer := GetTransformer(file.File, config)
			transformer.ToInjure()
			code := util.GetFileCode(file.File)
			println(code)
		}
	}
}

func TestLoadConfigration(t *testing.T) {
	LoadConfiguration()
}

func TestRandom(t *testing.T) {
	var a, b int
	for i := 0; i < 1000; i++ {
		intn := rand.Intn(2)
		if intn >= 1 {
			a++
		} else {
			b++
		}
	}
	println(a)
	println(b)
}

func TestGoCodeDrillEntry(t *testing.T) {
	env := env.CreateFaultEnv("/Users/misery/GolandProjects/code_fault/mmap", "/Users/misery/GolandProjects/code_fault/mmap1")
	env.ConditionInversedFault("util(1/5).myStruct(1/3).myFunc(1/2).myVariable | (3/4)...").
		ConditionInversedFault("util(1/5).myStruct(1/3).myFunc(1/2).myVariable")
	f := FaultPerformerFactory{}
	err := f.SetEnv(env).Run()
	if err != nil {
		return
	}
}
