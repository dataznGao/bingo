package double_write_scene

import (
	"github.com/dataznGao/bingo/scene/double_write_scene/dal"
	"testing"
)

func TestDoubleWrite(t *testing.T) {
	t1 := &dal.TestEn{Id: 9, Name: "admin"}
	DoubleWrite(t1)
}
