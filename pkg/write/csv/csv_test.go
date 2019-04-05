package csv

import (
	php2 "github.com/KingzCheung/ki18n/pkg/read/php"
	"testing"
)

func TestCsv_FileWrite(t *testing.T) {
	php := php2.NewPhp("../../../lang")
	dr := php.DirRead()
	c := &Csv{dr}
	c.FileWrite("test.php")
}
