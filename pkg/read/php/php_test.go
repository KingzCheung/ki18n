package php

import (
	"testing"
)

func TestPhp_DirRead(t *testing.T) {
	php := NewPhp("../../../lang")
	dr := php.DirRead()
	if len(dr) == 0 {
		t.Errorf("read php file and parse error!")
	}
}
