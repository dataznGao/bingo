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
	"log"
)

func init() {
	constant.InitFaultTypeMap()
}

func FillPackage(files map[string]*ds.File, notGoFiles []string) error {
	log.Printf("[bingo] INFO ====== 开始填充输出包 ======")
	for k, v := range files {
		newPath := util.CompareAndExchange(k, config.Config.OutputPath, config.Config.InputPath)
		if !v.IsInjured {
			log.Printf("[bingo] INFO 开始填充输出包, 文件: %v", k)
			err := util.CreateFile(newPath, util.GetFileCode(v.File))
			if err != nil {
				return err
			}
		}
		comment := util.GetBuildInfo(k)
		if len(comment) != 0 {
			log.Printf("[bingo] INFO 开始添加build信息, 文件: %v, \nbuild_info: %v", k, string(comment))
			err := util.InsertFileHead(newPath, comment)
			if err != nil {
				return err
			}
		}
	}
	for _, file := range notGoFiles {
		log.Printf("[bingo] INFO 开始填充输出包, 文件: %v", file)
		readFile, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		err = util.CreateFile(util.CompareAndExchange(file, config.Config.OutputPath, config.Config.InputPath),
			readFile)
		if err != nil {
			return err
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

// LoadPackage 加载需要变异的文件夹，返回文件名对应的文件，以及包对应的文件
func LoadPackage(path string) (map[string]*ds.File, []string, error) {
	m, notGoFiles, err := util.LoadAllFile(path)
	if err != nil {
		return nil, nil, err
	}
	files := make(map[string]*ds.File, 0)
	fset := token.NewFileSet() // positions are relative to fset
	for _, file := range m {
		f, err := parser.ParseFile(fset, file, nil, 0)
		if err != nil {
			return nil, nil, err
		}
		fi := &ds.File{
			File:       f,
			IsInjured:  false,
			FileName:   file,
			InputPath:  config.Config.InputPath,
			OutputPath: config.Config.OutputPath,
		}
		files[file] = fi
	}

	return files, notGoFiles, nil
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
