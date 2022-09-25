package env

import (
	"github.com/dataznGao/go_drill/config"
	"github.com/dataznGao/go_drill/constant"
	"log"
	"strings"
)

func CreateFaultEnv(inputPath, outputPath string) *FaultEnv {
	return &FaultEnv{
		InputPath:  inputPath,
		OutputPath: outputPath,
	}
}

type FaultEnv struct {
	InputPath   string
	OutputPath  string
	FaultPoints []*config.FaultConfig
}

// LocationPattern util(1/5).myStruct(1/3).myFunc(1/2).myVariable | main.*.*.*
type LocationPattern string

func (fe *FaultEnv) ValueFault(locationPattern LocationPattern, targetValue interface{}) *FaultEnv {
	lps, err := locationPattern.parse()
	if err != nil {
		log.Fatalf("[ValueFault] set fault point failed, err: %v", err.Error())
	}
	fe.FaultPoints = append(fe.FaultPoints, &config.FaultConfig{
		FaultType:        constant.StrValueFault,
		LocationPatterns: lps,
		TargetValue:      targetValue,
	})
	return fe
}

func (fe *FaultEnv) ConditionInversedFault(locationPattern LocationPattern) *FaultEnv {
	lps, err := locationPattern.parse()
	if err != nil {
		log.Fatalf("[ConditionInversedFault] set fault point failed, err: %v", err.Error())
	}
	fe.FaultPoints = append(fe.FaultPoints, &config.FaultConfig{
		FaultType:        constant.StrConditionInversedFault,
		LocationPatterns: lps,
		TargetValue:      nil,
	})
	return fe
}

func (fe *FaultEnv) SwitchMissDefaultFault(locationPattern LocationPattern) *FaultEnv {
	lps, err := locationPattern.parse()
	if err != nil {
		log.Fatalf("[ConditionInversedFault] set fault point failed, err: %v", err.Error())
	}
	fe.FaultPoints = append(fe.FaultPoints, &config.FaultConfig{
		FaultType:        constant.StrSwitchMissDefaultFault,
		LocationPatterns: lps,
	})
	return fe
}

func (fe *FaultEnv) NullFault(locationPattern LocationPattern) *FaultEnv {
	lps, err := locationPattern.parse()
	if err != nil {
		log.Fatalf("[ConditionInversedFault] set fault point failed, err: %v", err.Error())
	}
	fe.FaultPoints = append(fe.FaultPoints, &config.FaultConfig{
		FaultType:        constant.StrNullFault,
		LocationPatterns: lps,
		TargetValue:      "nil",
	})
	return fe
}

func (l LocationPattern) parse() ([]*config.LocationPatternP, error) {
	var (
		parts = make([]string, 0)
		res   = make([]*config.LocationPatternP, 0)
	)
	makeFaultPattern := func(sp string, p []*config.FaultPatternP) []*config.FaultPatternP {
		if !strings.Contains(sp, "(") {
			p = append(p, &config.FaultPatternP{
				Name:           strings.TrimSpace(sp),
				ActivationRate: "1",
			})
		} else {
			faultPatterns := strings.Split(strings.TrimSpace(sp), "(")
			p = append(p, &config.FaultPatternP{
				Name:           faultPatterns[0],
				ActivationRate: strings.TrimRight(faultPatterns[1], ")"),
			})
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
		lp := &config.LocationPatternP{
			PackageP:  nil,
			StructP:   nil,
			MethodP:   nil,
			VariableP: nil,
		}
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

func (fe *FaultEnv) ExceptionUncaughtFault(locationPattern LocationPattern) *FaultEnv {
	lps, err := locationPattern.parse()
	if err != nil {
		log.Fatalf("[ExceptionUncaughtFault] set fault point failed, err: %v", err.Error())
	}
	fe.FaultPoints = append(fe.FaultPoints, &config.FaultConfig{
		FaultType:        constant.StrExceptionUncaughtFault,
		LocationPatterns: lps,
		TargetValue:      nil,
	})
	return fe
}

func (fe *FaultEnv) ExceptionUnhandledFault(locationPattern LocationPattern) *FaultEnv {
	lps, err := locationPattern.parse()
	if err != nil {
		log.Fatalf("[ExceptionUncaughtFault] set fault point failed, err: %v", err.Error())
	}
	fe.FaultPoints = append(fe.FaultPoints, &config.FaultConfig{
		FaultType:        constant.StrExceptionUnhandledFault,
		LocationPatterns: lps,
		TargetValue:      nil,
	})
	return fe
}

func (fe *FaultEnv) ExceptionShortcircuitFault(locationPattern LocationPattern) *FaultEnv {
	lps, err := locationPattern.parse()
	if err != nil {
		log.Fatalf("[ExceptionUncaughtFault] set fault point failed, err: %v", err.Error())
	}
	fe.FaultPoints = append(fe.FaultPoints, &config.FaultConfig{
		FaultType:        constant.StrExceptionUnhandledFault,
		LocationPatterns: lps,
		TargetValue:      nil,
	})
	return fe
}

func (fe *FaultEnv) SyncFault(locationPattern LocationPattern) *FaultEnv {
	lps, err := locationPattern.parse()
	if err != nil {
		log.Fatalf("[ExceptionUncaughtFault] set fault point failed, err: %v", err.Error())
	}
	fe.FaultPoints = append(fe.FaultPoints, &config.FaultConfig{
		FaultType:        constant.StrSyncFault,
		LocationPatterns: lps,
		TargetValue:      nil,
	})
	return fe
}

func (fe *FaultEnv) AttributeReversoFault(locationPattern LocationPattern, targetValue interface{}) *FaultEnv {
	lps, err := locationPattern.parse()
	if err != nil {
		log.Fatalf("[ValueFault] set fault point failed, err: %v", err.Error())
	}
	fe.FaultPoints = append(fe.FaultPoints, &config.FaultConfig{
		FaultType:        constant.StrAttributeReversoFault,
		LocationPatterns: lps,
		TargetValue:      targetValue,
	})
	return fe
}
