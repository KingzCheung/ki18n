package csv

import (
	csv2 "github.com/KingzCheung/ki18n/pkg/read/csv"
	php2 "github.com/KingzCheung/ki18n/pkg/read/php"
	"github.com/KingzCheung/ki18n/pkg/util"
	"testing"
)

func TestCsv_FileWrite(t *testing.T) {
	php := php2.NewPhp("../../../lang")
	dr := php.DirRead()
	c := &Csv{dr}
	testFile := "test.csv"
	c.FileWrite(testFile)
	if b, _ := util.PathExists(testFile); !b {
		t.Errorf("FileWrite error,create csv fail.")
	}

	csv := csv2.NewCsv(testFile)
	r := csv.Read()
	if len(r) == 0 {
		t.Errorf("parse csv fail")
	}
}
