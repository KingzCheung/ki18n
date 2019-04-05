package csv

import (
	"fmt"
	"github.com/gookit/color"
	"os"
)

type Csv struct {
	content []map[string]string
}

func (c *Csv) FileWrite(path string) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		color.Red.Println(err)
	}
	_, _ = f.WriteString("KEY,简体中文,繁体中文,英文,备注")

	len := len(c.content)

	var csv [][]string
	for _,m := range c.content {
		for i := 0; i < len; i++ {
			csv = append(csv, c.content[i][])
		}
	}

	fmt.Println(csv)
}
