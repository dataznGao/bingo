package code_drill

import (
	"github.com/dataznGao/go_drill/config"
	"github.com/dataznGao/go_drill/constant"
	"github.com/dataznGao/go_drill/ds"
	"github.com/dataznGao/go_drill/env"
	"github.com/dataznGao/go_drill/transformer"
	"github.com/dataznGao/go_drill/transformer/condition_inversed"
	"github.com/dataznGao/go_drill/transformer/exception_uncaught"
	"github.com/dataznGao/go_drill/transformer/exception_unhandled"
	"github.com/dataznGao/go_drill/transformer/reverso"
	"github.com/dataznGao/go_drill/transformer/switch_miss_default"
	"github.com/dataznGao/go_drill/transformer/sync"
	"github.com/dataznGao/go_drill/transformer/value"
	"github.com/dataznGao/go_drill/util"
	"github.com/jinzhu/copier"
	"go/ast"
	"go/parser"
	"go/token"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var Config *config.Configuration

func init() {
	constant.InitFaultTypeMap()
}

type FaultPerformerFactory struct {
	_env *env.FaultEnv
}

func (f *FaultPerformerFactory) SetEnv(env *env.FaultEnv) *FaultPerformerFactory {
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

func goCodeYamlDrillEntry() error {
	conf, err := LoadConfiguration()
	if err != nil {
		log.Fatalf("[LoadConfiguration] config file %v is wrong", constant.ConfigFile)
		return err
	}
	return entry(conf)
}

func goCodeDrillEntry(env *env.FaultEnv) error {
	conf := &config.Configuration{
		InputPath:   env.InputPath,
		OutputPath:  env.OutputPath,
		FaultPoints: env.FaultPoints,
	}
	return entry(conf)
}

func entry(config *config.Configuration) error {
	var originInputPath = config.InputPath
	Config = config
	var err error
	files, err := LoadPackage(Config.InputPath)
	if err != nil {
		return err
	}
	gc := make([]string, 0)
	for filename, file := range files {
		for _, faultConfig := range Config.FaultPoints {
			err, newPath, replica, _ := PerformInjure(file, filename, faultConfig)
			gc = append(gc, replica)
			if err != nil {
				log.Fatalf("[PerformInjure] file: %v injure fault failed, err: %v", filename, err.Error())
			}
			err = fillPackage(files)
			if err != nil {
				return err
			}
			command := "cd " + config.OutputPath + " && go build"

			_, err = util.Command(command)
			if err != nil {
				util.Copy(replica, newPath)
				fset := token.NewFileSet()
				parseFile, _ := parser.ParseFile(fset, newPath, nil, 0)
				file.File = parseFile
			}
		}
	}
	config.InputPath = originInputPath
	return util.Clean(gc)
}

func fillPackage(files map[string]*ds.FileInjure) error {
	for k, v := range files {
		if !v.IsInjured {
			err := util.CreateFile(util.CompareAndExchange(k, Config.OutputPath, Config.InputPath),
				util.GetFileCode(v.File))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func PerformInjure(fi *ds.FileInjure, fileName string, config *config.FaultConfig) (error, string, string, *ast.File) {
	faultTransformer := GetTransformer(fi.File, config)
	if faultTransformer == nil {
		return constant.NewNoFaultTypeError(config.FaultType), "", "", nil
	}

	replicaFile := new(ast.File)
	copier.Copy(&replicaFile, &fi.File)

	faultTransformer.ToInjure()
	code := util.GetFileCode(fi.File)

	newPath := util.CompareAndExchange(fileName, Config.OutputPath, Config.InputPath)
	replica := newPath + "1"
	util.Copy(newPath, replica)
	err := util.CreateFile(newPath, code)
	fi.IsInjured = true
	if err != nil {
		return err, "", "", nil
	}
	return nil, newPath, replica, replicaFile
}

func LoadConfiguration() (*config.Configuration, error) {
	configFile, err := ioutil.ReadFile(constant.ConfigFile)
	if err != nil {
		return nil, err
	}
	conf := &config.Configuration{}
	err = yaml.Unmarshal(configFile, Config)
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

func GetTransformer(file *ast.File, config *config.FaultConfig) transformer.Transformer {
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
