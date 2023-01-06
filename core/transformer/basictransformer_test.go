package transformer

import (
	"github.com/dataznGao/bingo/util"
	"testing"
)

func TestCreateFile(t *testing.T) {
	newPath := "/Users/misery/GolandProjects/bingo/scene1/double_write_scene/rpc/helloservice.go"

	command := "cd " + util.GetFather(newPath) + " && go build"
	_, err := util.Command(command)
	if err != nil {

	}
}
