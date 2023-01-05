package core

import (
	"github.com/dataznGao/bingo/constant"
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/core/ds"
	"github.com/dataznGao/bingo/core/transformer"
	"github.com/dataznGao/bingo/core/transformer/condition_inversed"
	"github.com/dataznGao/bingo/core/transformer/exception_uncaught"
	"github.com/dataznGao/bingo/core/transformer/exception_unhandled"
	"github.com/dataznGao/bingo/core/transformer/reverso"
	"github.com/dataznGao/bingo/core/transformer/switch_miss_default"
	"github.com/dataznGao/bingo/core/transformer/sync"
	"github.com/dataznGao/bingo/core/transformer/value"
	"github.com/dataznGao/bingo/util"
	"go/parser"
	"go/token"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func init() {
	constant.InitFaultTypeMap()
}

func FillPackage(files map[string]*ds.FileInjure) error {
	for k, v := range files {
		if !v.IsInjured {
			err := util.CreateFile(util.CompareAndExchange(k, config.Config.OutputPath, config.Config.InputPath),
				util.GetFileCode(v.File))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func PerformInjure(fi *ds.File, conf *config.FaultConfig) error {
	faultTransformer := getTransformer(fi, conf)
	if faultTransformer == nil {
		return constant.NewNoFaultTypeError(conf.FaultType)
	}
	faultTransformer.ToInjure()
	return nil
}

func loadConfiguration() (*config.Configuration, error) {
	configFile, err := ioutil.ReadFile(constant.ConfigFile)
	if err != nil {
		return nil, err
	}
	conf := &config.Configuration{}
	err = yaml.Unmarshal(configFile, config.Config)
	return conf, nil
}

// LoadPackage 加载需要注入的文件夹，返回文件名对应的文件，以及包对应的文件
func LoadPackage(path string) (map[string]*ds.FileInjure, error) {
	m, err := util.LoadAllGoFile(path)
	if err != nil {
		return nil, err
	}
	files := make(map[string]*ds.FileInjure, 0)
	fset := token.NewFileSet() // positions are relative to fset
	for _, file := range m {
		f, err := parser.ParseFile(fset, file, nil, 0)
		if err != nil {
			return nil, err
		}
		fi := &ds.FileInjure{
			File:      f,
			IsInjured: false,
		}
		files[file] = fi
	}

	return files, nil
}

func getTransformer(file *ds.File, config *config.FaultConfig) transformer.Transformer {
	switch constant.FaultTypeMap[config.FaultType] {
	case constant.ConditionInversedFault:
		return &condition_inversed.ConditionInversedTransformer{
			File:   file,
			Config: config,
		}
	case constant.ValueFault:
		return &value.ValueTransformer{
			File:   file,
			Config: config,
		}
	case constant.NullFault:
		return &value.ValueTransformer{
			File:   file,
			Config: config,
		}
	case constant.SwitchMissDefaultFault:
		return &switch_miss_default.SwitchMissDefaultTransformer{
			File:   file,
			Config: config,
		}
	case constant.ExceptionUnhandledFault:
		return &exception_unhandled.ExceptionUnhandledTransformer{
			File:   file,
			Config: config,
		}
	case constant.ExceptionUncaughtFault:
		return &exception_uncaught.ExceptionUncaughtTransformer{
			File:   file,
			Config: config,
		}
	case constant.ExceptionShortcircuitFault:
		return &exception_unhandled.ExceptionUnhandledTransformer{
			File:   file,
			Config: config,
		}
	case constant.SyncFault:
		return &sync.SyncTransformer{
			File:   file,
			Config: config,
		}
	case constant.AttributeReversoFault:
		return &reverso.AttributeReversoTransformer{
			File:   file,
			Config: config,
		}
	}
	return nil
}
