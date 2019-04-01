package csv

import (
	"fmt"
	"testing"
)

func TestCsv_Read(t *testing.T) {
	csv := NewCsv("../../../language.csv")
	csvSlice := csv.Read()
	fmt.Println(csvSlice)
}
