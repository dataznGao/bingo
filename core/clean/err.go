package clean

import (
	"github.com/dataznGao/bingo/constant"
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/util"
	"os"
	"strconv"
	"strings"
)

// # github.com/douyu/jupiter/pkg/executor/xxl
//./task.go:76:7: undefined: err
//./task.go:79:52: undefined: err
type ErrCleaner struct {
	Err string
	// 错误类型以及文件名和错误位置
	position map[string]map[ErrType][]*Pos
	// 文件位置
	FileName string
}

type ErrType int

const IMPORT_AND_NOT_USED ErrType = 1
const DECLARED_BUT_NOT_USED ErrType = 2
const ERR_UNDEFINED ErrType = 3
const NO_NEW_VARIABLES ErrType = 4
const REDECLARE ErrType = 5

type Pos struct {
	FileName string
	Line     int
	Extra    string
}

func (t *ErrCleaner) ConvertErrMsg() {
	t.position = make(map[string]map[ErrType][]*Pos)

	split := strings.Split(t.Err, "\n")
	prefix := ""
	packageName := util.GetPackageName(config.Config.OutputPath)
	for _, s := range split {
		s = strings.TrimSpace(s)
		// 说明是错误文件的开头
		if strings.HasPrefix(s, "# ") {
			prefix = strings.ReplaceAll(s, "# ", "")
		} else if strings.HasPrefix(s, ".") {
			newPrefix := prefix
			if strings.Contains(s, "../") {
				cnt := strings.Count(s, "../")
				s = strings.ReplaceAll(s, "../", "")
				if !strings.HasPrefix(s, "./") {
					s = "./" + s
				}
				for i := 0; i < cnt; i++ {
					n := len(newPrefix)
					for index := n - 1; index >= 0; index-- {
						if newPrefix[index] == '/' {
							newPrefix = newPrefix[:index]
							break
						}
					}
				}
			}
			// 文件名:行号:列号: 错误详情
			part := strings.Split(s[1:], ":")
			if len(part) < 4 {
				continue
			}
			extra := ""
			for i := 3; i < len(part); i++ {
				extra += ":" + part[i]
			}
			extra = extra[1:]
			part[0] = strings.TrimSpace(part[0])
			pos := newPrefix + part[0]
			line := part[1]
			atoi, _ := strconv.Atoi(line)
			// 对错误类型进行枚举, 并且给出详情错误
			errType, extra := decideErrTypeByStr(extra)
			filePath := strings.Replace(pos, packageName, config.Config.OutputPath, 1)

			if t.position[filePath] == nil {
				t.position[filePath] = make(map[ErrType][]*Pos, 0)
			}
			if t.position[filePath][errType] == nil {
				t.position[filePath][errType] = make([]*Pos, 0)
			}
			t.position[filePath][errType] = append(t.position[filePath][errType], &Pos{
				FileName: filePath,
				Line:     atoi,
				Extra:    strings.TrimSpace(extra),
			})

		} else {
			n := len(prefix)
			newPrefix := prefix

			for i := 0; i < n; i++ {
				if newPrefix[i] == '/' {
					newPrefix = newPrefix[i+1:]
					break
				}
			}
			if !strings.HasPrefix(s, "./") {
				s = "./" + s
			}
			// 文件名:行号:列号: 错误详情
			part := strings.Split(s[1:], ":")
			if len(part) < 4 {
				continue
			}
			extra := ""
			for i := 3; i < len(part); i++ {
				extra += ":" + part[i]
			}
			extra = extra[1:]
			part[0] = strings.TrimSpace(part[0])
			pos := newPrefix + part[0]
			line := part[1]
			atoi, _ := strconv.Atoi(line)
			// 对错误类型进行枚举, 并且给出详情错误
			errType, extra := decideErrTypeByStr(extra)
			filePath := strings.Replace(pos, packageName, config.Config.OutputPath, 1)

			if t.position[filePath] == nil {
				t.position[filePath] = make(map[ErrType][]*Pos, 0)
			}
			if t.position[filePath][errType] == nil {
				t.position[filePath][errType] = make([]*Pos, 0)
			}
			t.position[filePath][errType] = append(t.position[filePath][errType], &Pos{
				FileName: filePath,
				Line:     atoi,
				Extra:    strings.TrimSpace(extra),
			})
		}
	}

}

func (t *ErrCleaner) Fix() error {
	t.ConvertErrMsg()
	err := t.OneFix()
	if err != nil {
		return err
	}
	command := "cd " + util.GetFather(t.FileName) + " && go build"
	res, err := util.Command(command)
	if err != nil {
		t.Err = res
		return t.Fix()
	}
	return nil
}

// OneFix 需要以文件为粒度，内部进行错误的划分
func (t *ErrCleaner) OneFix() error {
	var err error
	for fileName, m := range t.position {
		for errType, pos := range m {
			// 准备好输入数组
			if _, ok := t.position[fileName][0]; ok {
				return constant.NewCannotSolveError()
			}
			contentWithLine := make([]string, 0)
			con, err := os.ReadFile(fileName)
			if err != nil {
				return err
			}
			var content = string(con)
			split := strings.Split(content, "\n")
			for _, s := range split {
				contentWithLine = append(contentWithLine, s)
			}
			// 修bug
			contentWithLine, err = fix(contentWithLine, errType, pos)
			if err != nil {
				return err
			}
			// fmt
			newContent := ""
			for _, i := range contentWithLine {
				newContent += i + "\n"
			}
			err = util.CreateFile(fileName, []byte(newContent))
			if err != nil {
				return err
			}
			_, err = util.Command("gofmt -w " + fileName)
			if err != nil {
				return err
			}
		}
	}
	return err
}

// 对pos进行不同的errType进行修理
func fix(contentWithLine []string, errType ErrType, pos []*Pos) ([]string, error) {
	for _, po := range pos {
		if len(contentWithLine) <= po.Line-1 {
			return contentWithLine, nil
		}
		if errType == IMPORT_AND_NOT_USED {
			if strings.Contains(contentWithLine[po.Line-1], po.Extra) {
				contentWithLine[po.Line-1] = ""
			}
		} else if errType == DECLARED_BUT_NOT_USED {
			if strings.Contains(contentWithLine[po.Line-1], po.Extra) {
				contentWithLine[po.Line-1] = strings.ReplaceAll(contentWithLine[po.Line-1], po.Extra, "_")
			}
		} else if errType == ERR_UNDEFINED {
			if po.Line-1 > 0 {
				bf := contentWithLine[po.Line-2]
				declare := "var " + po.Extra + " error"
				if !strings.Contains(bf, declare) {
					contentWithLine = append(contentWithLine[:po.Line-1], append([]string{declare}, contentWithLine[po.Line-1:]...)...)
				}
			}
		} else if errType == NO_NEW_VARIABLES {
			str := contentWithLine[po.Line-1]
			str = strings.ReplaceAll(str, ":=", "=")
			contentWithLine[po.Line-1] = str
		} else if errType == REDECLARE {
			contentWithLine[po.Line-1] = strings.ReplaceAll(contentWithLine[po.Line-1], po.Extra, "_")
		}
	}
	return contentWithLine, nil
}

func decideErrTypeByStr(str string) (ErrType, string) {
	if strings.Contains(str, "imported and not used:") {
		extra := strings.ReplaceAll(str, "imported and not used:", "")
		if strings.Contains(extra, " as ") {
			extra = strings.Split(extra, " as ")[0]
		}
		extra = strings.TrimSpace(extra)
		extra = strings.TrimRight(extra, "\"")
		extra = strings.TrimLeft(extra, "\"")
		return IMPORT_AND_NOT_USED, extra
	} else if strings.Contains(str, "declared but not used") {
		extra := strings.ReplaceAll(str, " declared but not used", "")
		return DECLARED_BUT_NOT_USED, extra
	} else if strings.Contains(str, "undefined: ") {
		if strings.Contains(str, "err") {
			extra := strings.ReplaceAll(str, "undefined: ", "")
			return ERR_UNDEFINED, strings.TrimSpace(extra)
		}
	} else if strings.Contains(str, "no new variables on left side of :=") {
		return NO_NEW_VARIABLES, ""
	} else if strings.Contains(str, "redeclared in this block") {
		return REDECLARE, strings.ReplaceAll(str, " redeclared in this block", "")
	}
	return 0, str
}
