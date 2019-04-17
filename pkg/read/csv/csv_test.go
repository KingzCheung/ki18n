package csv

import (
	"testing"
)

func TestCsv_Read(t *testing.T) {
	csv := NewCsv("../../../language.csv")
	csvSlice := csv.Read()
	if len(csvSlice) == 0 {
		t.Errorf("数组的长度不应该是0")
	}
}
