package clean

import (
	"github.com/dataznGao/bingo/core/config"
	"testing"
)

func TestErrCleaner_ConvertErrMsg(t *testing.T) {
	config.Config = &config.Configuration{
		InputPath: "/Users/misery/GolandProjects/tmp_enhance",
	}
	cleaner := &ErrCleaner{
		Err: "# github.com/douyu/jupiter/pkg/conf\n./datasource.go:30:5: undefined: err\n./datasource.go:31:15: undefined: err\n",
	}
	cleaner.ConvertErrMsg()
	cleaner.Fix()
}
