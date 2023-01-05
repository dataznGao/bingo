package bingo

import (
	"github.com/dataznGao/bingo/core"
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/core/ds"
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
	files, err := core.LoadPackage(conf.InputPath)
	if err != nil {
		return err
	}
	for filename, file := range files {
		for _, faultConfig := range conf.FaultPoints {
			err := core.PerformInjure(&ds.File{
				File:       file.File,
				FileName:   filename,
				InputPath:  conf.InputPath,
				OutputPath: conf.OutputPath,
			}, faultConfig)
			if err != nil {
				log.Fatalf("[performInjure] file: %v injure fault failed, err: %v", filename, err.Error())
			}
			err = core.FillPackage(files)
			if err != nil {
				return err
			}

		}
	}
	return nil
}
