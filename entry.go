package bingo

import (
	"github.com/dataznGao/bingo/constant"
	"github.com/dataznGao/bingo/core"
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/core/run-test"
	"github.com/dataznGao/bingo/task"
	"github.com/dataznGao/bingo/util"
	"log"
	"time"
)

type MutationPerformer struct {
	_env *MutationEnv
}

func (f *MutationPerformer) SetEnv(env *MutationEnv) *MutationPerformer {
	return &MutationPerformer{
		_env: env,
	}
}

// Run 仅仅对inputPath中的进行变异，不进行测试
func (f *MutationPerformer) Run() error {
	err := bingoMutationEntry(f._env)
	if err != nil {
		return err
	}
	return nil
}

// Test 对inputPath中的进行变异测试，返回两次测试的结果对比
func (f *MutationPerformer) Test() error {
	err := f.Run()
	if err != nil {
		return err
	}
	group := task.NewGroup(2)
	res := make([]string, 2)
	// 1. 对变异前进行代码测试
	task1 := func() {
		log.Printf("[bingo] INFO 变异前代码测试")
		rawResult, err := run.Test(f._env.InputTestPath, f._env.InputPath)
		if err != nil {
			log.Fatalf("[bingo] ERROR 变异前代码测试失败, err: %v", err)
		}
		log.Printf("[bingo] INFO 变异前代码测试完毕")
		res[0] = rawResult
	}
	group.Add(task1)
	// 2. 对变异后的代码进行测试
	task2 := func() {
		log.Printf("[bingo] INFO 变异后代码测试")
		mutResult, err := run.Test(f._env.OutputTestPath, f._env.OutputPath)
		if err != nil {
			log.Fatalf("[bingo] ERROR 变异后代码测试失败, err: %v", err)
		}
		log.Printf("[bingo] INFO 变异后代码测试完毕")
		res[1] = mutResult
	}
	group.Add(task2)
	group.Start()
	group.Wait()
	// 报告路径
	reportPath := f._env.InputPath + constant.Separator + "bingo"
	fileName := util.Time2Str(time.Now())
	result := "变异测试前：\n" + res[0] + "\n" + "变异测试后：\n" + res[1]
	return util.CreateFile(reportPath+constant.Separator+fileName, []byte(result))
}

// 官方唯一默认指定入口
func bingoMutationEntry(env *MutationEnv) error {
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
	err = core.FillPackage(files, notGoFiles)
	if err != nil {
		return err
	}
	log.Printf("[bingo] INFO ======= 开始变异 ========")
	for filename, file := range files {
		for _, faultConfig := range conf.FaultPoints {
			log.Printf("[bingo] INFO 本次变异文件: %v, 变异算子: %v", filename, faultConfig.FaultType)
			err := core.PerformInjure(file, faultConfig)
			if err != nil {
				log.Fatalf("[bingo] FATAL file: %v injure fault failed, err: %v", filename, err.Error())
			}
		}
	}

	return nil
}
