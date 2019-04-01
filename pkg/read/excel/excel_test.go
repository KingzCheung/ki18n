package excel

import "testing"

func TestExcel_Read(t *testing.T) {
	excel := NewExcel("../../../language.xlsx")
	excel.Read()
}
