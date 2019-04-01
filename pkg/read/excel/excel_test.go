package excel

import (
	"testing"
)

func TestExcel_Read(t *testing.T) {
	excel := NewExcel("../../../language.xlsx")
	r := excel.Read()
	if len(r) == 0 {
		t.Errorf("找不到内容")
	}
}
