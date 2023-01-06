package bingo

import (
	"github.com/dataznGao/bingo/core"
	"github.com/dataznGao/bingo/core/config"
	"log"
)

type FaultPerformerFactory struct {
	_env *FaultEnv
}

func (f *FaultPerformerFactory) SetEnv(env *FaultEnv) *FaultPerformerFactory {
	return &FaultPerformerFactory{
		_env: env,
	}
}

func (f *FaultPerformerFactory) Run() error {
	err := goCodeDrillEntry(f._env)
	if err != nil {
		return err
	}
	return nil
}

func goCodeDrillEntry(env *FaultEnv) error {
	conf := &config.Configuration{
		InputPath:   env.InputPath,
		OutputPath:  env.OutputPath,
		FaultPoints: env.FaultPoints,
	}
	return entry(conf)
}

func entry(conf *config.Configuration) error {
	config.Config = conf
	var err error
	files, notGoFiles, err := core.LoadPackage(conf.InputPath)
	if err != nil {
		return err
	}
	log.Printf("[bingo] INFO ======= 开始变异 ========")
	for filename, file := range files {
		for _, faultConfig := range conf.FaultPoints {
			log.Printf("[bingo] INFO 本次变异算子: %v", faultConfig.FaultType)
			err := core.PerformInjure(file, faultConfig)
			if err != nil {
				log.Fatalf("[performInjure] file: %v injure fault failed, err: %v", filename, err.Error())
			}
		}
	}
	err = core.FillPackage(files, notGoFiles)
	if err != nil {
		return err
	}
	return nil
}
