package util

import (
	"github.com/dataznGao/go_drill/config"
	"github.com/dataznGao/go_drill/constant"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func CompareAndExchange(oldPath, newPath, inputPath string) string {
	if strings.HasPrefix(oldPath, inputPath) {
		replace := strings.Replace(oldPath, inputPath, newPath, 1)
		return replace
	}
	return ""
}

func CompareAndExchangeTestPath(oldPath, newPath, inputPath string) string {
	if strings.HasPrefix(oldPath, inputPath) {
		replace := strings.Replace(oldPath, inputPath, newPath, 1)
		return replace
	}
	return ""
}

func Contains(elem string, elems []string) bool {
	if elem == "*" {
		return true
	}
	for _, s := range elems {
		if s == elem {
			return true
		}
	}
	return false
}

func ShowLocatePackage(elem string, lp []*config.LocationPatternP) []*config.LocationPatternP {
	res := make([]*config.LocationPatternP, 0)
	elem = strings.TrimSpace(elem)
	for _, str := range lp {
		packageName := strings.TrimSpace(str.PackageP.Name)
		if elem == packageName || packageName == "*" || packageName == "" {
			res = append(res, str)
		}
	}
	return res
}

func ShowLocateStruct(elem string, lp []*config.LocationPatternP) []*config.LocationPatternP {
	res := make([]*config.LocationPatternP, 0)
	elem = strings.TrimSpace(elem)
	for _, str := range lp {
		structName := strings.TrimSpace(str.StructP.Name)
		if elem == structName || structName == "*" || structName == "" {
			res = append(res, str)
		}
	}
	return res
}

func ShowLocateMethod(elem string, lp []*config.LocationPatternP) []*config.LocationPatternP {
	res := make([]*config.LocationPatternP, 0)
	elem = strings.TrimSpace(elem)
	for _, str := range lp {
		methodName := strings.TrimSpace(str.MethodP.Name)
		if elem == methodName || methodName == "*" || methodName == "" {
			res = append(res, str)
		}
	}
	return res
}

func ShowLocateVariable(elem string, lp []*config.LocationPatternP) []*config.LocationPatternP {
	res := make([]*config.LocationPatternP, 0)
	elem = strings.TrimSpace(elem)
	for _, str := range lp {
		variable := strings.TrimSpace(str.VariableP.Name)
		if elem == variable || variable == "*" || variable == "" {
			res = append(res, str)
		}
	}
	return res
}

func CanPerform(p string) bool {
	var (
		err error
		a   int
		b   int
	)
	if strings.TrimSpace(p) == "1" {
		return true
	}
	rand.Seed(time.Now().Unix())
	split := strings.Split(p, "/")
	if len(split) != 2 {
		log.Fatal(constant.NewProbabilityError(p))
		return false
	}
	a, err = strconv.Atoi(strings.TrimSpace(split[0]))
	if err != nil {
		log.Fatal(constant.NewProbabilityError(p))
		return false
	}
	b, err = strconv.Atoi(strings.TrimSpace(split[1]))
	if err != nil {
		log.Fatal(constant.NewProbabilityError(p))
		return false
	}
	intn := rand.Intn(b)
	if intn >= a {
		return false
	} else {
		return true
	}
}
