package php

import (
	"fmt"
	"testing"
)

func TestPhp_DirRead(t *testing.T) {
	php := NewPhp("../../../lang")
	dr := php.DirRead()
	fmt.Println(dr)
}
