package util

import "testing"

func TestEscape(t *testing.T) {
	str := `"hello"`
	rstr := Escape(str)
	rightRes := `\"hello\"`
	if rstr != rightRes {
		t.Errorf("Escape 函数出错：本来应该是 (%s),结果却是:%s", rightRes, rstr)
	}
}
