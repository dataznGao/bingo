package dal

import "testing"

func TestAddTest(t *testing.T) {
	test := &TestEn{Id: 0, Name: "binjie"}
	err := AddTest(test)
	if err != nil {
		return
	}
	return
}
