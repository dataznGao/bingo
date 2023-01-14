package bingo

import (
	"github.com/dataznGao/bingo/constant"
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/util"
	"log"
	"strings"
)

func CreateMutationEnv(inputPath, outputPath, testPath string) *MutationEnv {
	if strings.HasSuffix(outputPath, constant.Separator) {
		outputPath = outputPath[:len(outputPath)-1]
	}
	outputTestPath := util.CompareAndExchange(testPath, outputPath, inputPath)
	return &MutationEnv{InputPath: inputPath, OutputPath: outputPath, InputTestPath: testPath, OutputTestPath: outputTestPath}
}

type MutationEnv struct {
	InputPath      string
	OutputPath     string
	InputTestPath  string
	OutputTestPath string
	FaultPoints    []*config.FaultConfig
}
type LocationPattern string

func (fe *MutationEnv) ValueFault(locationPattern LocationPattern, targetValue interface{}) *MutationEnv {
	lps, err := locationPattern.parse()
	if err != nil {
		log.Fatalf("[ValueFault] set fault point failed, err: %v", err.Error())
	}
	fe.FaultPoints = append(fe.FaultPoints, &config.FaultConfig{FaultType: constant.StrValueFault, LocationPatterns: lps, TargetValue: targetValue})
	return fe
}
func (fe *MutationEnv) ConditionInversedFault(locationPattern LocationPattern) *MutationEnv {
	lps, err := locationPattern.parse()
	if err != nil {
		log.Fatalf("[ConditionInversedFault] set fault point failed, err: %v", err.Error())
	}
	fe.FaultPoints = append(fe.FaultPoints, &config.FaultConfig{FaultType: constant.StrConditionInversedFault, LocationPatterns: lps, TargetValue: nil})
	return fe
}
func (fe *MutationEnv) SwitchMissDefaultFault(locationPattern LocationPattern) *MutationEnv {
	lps, err := locationPattern.parse()
	if err != nil {
		log.Fatalf("[ConditionInversedFault] set fault point failed, err: %v", err.Error())
	}
	fe.FaultPoints = append(fe.FaultPoints, &config.FaultConfig{FaultType: constant.StrSwitchMissDefaultFault, LocationPatterns: lps})
	return fe
}
func (fe *MutationEnv) NullFault(locationPattern LocationPattern) *MutationEnv {
	lps, err := locationPattern.parse()
	if err != nil {
		log.Fatalf("[ConditionInversedFault] set fault point failed, err: %v", err.Error())
	}
	fe.FaultPoints = append(fe.FaultPoints, &config.FaultConfig{FaultType: constant.StrNullFault, LocationPatterns: lps, TargetValue: "nil"})
	return fe
}
func (l LocationPattern) parse() ([]*config.LocationPatternP, error) {
	var (
		parts = make([]string, 0)
		res   = make([]*config.LocationPatternP, 0)
	)
	makeFaultPattern := func(sp string, p []*config.FaultPatternP) []*config.FaultPatternP {
		if !strings.Contains(sp, "(") {
			p = append(p, &config.FaultPatternP{Name: strings.TrimSpace(sp), ActivationRate: "1"})
		} else {
			faultPatterns := strings.Split(strings.TrimSpace(sp), "(")
			p = append(p, &config.FaultPatternP{Name: faultPatterns[0], ActivationRate: strings.TrimRight(faultPatterns[1], ")")})
		}
		return p
	}
	if !strings.Contains(string(l), "|") {
		parts = append(parts, string(l))
	} else {
		parts = strings.Split(string(l), "|")
	}
	for _, part := range parts {
		split := strings.Split(part, ".")
		if len(split) != 4 {
			return nil, constant.NewLocationPatternFormatError(string(l))
		}
		lp := &config.LocationPatternP{PackageP: nil, StructP: nil, MethodP: nil, VariableP: nil}
		p := make([]*config.FaultPatternP, 0)
		for _, sp := range split {
			p = makeFaultPattern(sp, p)
		}
		if len(p) != 4 {
			return nil, constant.NewLocationPatternFormatError(string(l))
		}
		lp.PackageP = p[0]
		lp.StructP = p[1]
		lp.MethodP = p[2]
		lp.VariableP = p[3]
		res = append(res, lp)
	}
	return res, nil
}
func (fe *MutationEnv) ExceptionUncaughtFault(locationPattern LocationPattern) *MutationEnv {
	lps, err := locationPattern.parse()
	if err != nil {
		log.Fatalf("[ExceptionUncaughtFault] set fault point failed, err: %v", err.Error())
	}
	fe.FaultPoints = append(fe.FaultPoints, &config.FaultConfig{FaultType: constant.StrExceptionUncaughtFault, LocationPatterns: lps, TargetValue: nil})
	return fe
}
func (fe *MutationEnv) ExceptionUnhandledFault(locationPattern LocationPattern) *MutationEnv {
	lps, err := locationPattern.parse()
	if err != nil {
		log.Fatalf("[ExceptionUncaughtFault] set fault point failed, err: %v", err.Error())
	}
	fe.FaultPoints = append(fe.FaultPoints, &config.FaultConfig{FaultType: constant.StrExceptionUnhandledFault, LocationPatterns: lps, TargetValue: nil})
	return fe
}
func (fe *MutationEnv) ExceptionShortcircuitFault(locationPattern LocationPattern) *MutationEnv {
	lps, err := locationPattern.parse()
	if err != nil {
		log.Fatalf("[ExceptionUncaughtFault] set fault point failed, err: %v", err.Error())
	}
	fe.FaultPoints = append(fe.FaultPoints, &config.FaultConfig{FaultType: constant.StrExceptionUnhandledFault, LocationPatterns: lps, TargetValue: nil})
	return fe
}
func (fe *MutationEnv) SyncFault(locationPattern LocationPattern) *MutationEnv {
	lps, err := locationPattern.parse()
	if err != nil {
		log.Fatalf("[ExceptionUncaughtFault] set fault point failed, err: %v", err.Error())
	}
	fe.FaultPoints = append(fe.FaultPoints, &config.FaultConfig{FaultType: constant.StrSyncFault, LocationPatterns: lps, TargetValue: nil})
	return fe
}
func (fe *MutationEnv) AttributeReversoFault(locationPattern LocationPattern, targetValue interface{}) *MutationEnv {
	lps, err := locationPattern.parse()
	if err != nil {
		log.Fatalf("[ValueFault] set fault point failed, err: %v", err.Error())
	}
	fe.FaultPoints = append(fe.FaultPoints, &config.FaultConfig{FaultType: constant.StrAttributeReversoFault, LocationPatterns: lps, TargetValue: targetValue})
	return fe
}
