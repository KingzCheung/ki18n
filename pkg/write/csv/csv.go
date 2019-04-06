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
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		color.Red.Println(err)
	}

	defer f.Close()

	_, _ = f.WriteString("KEY,简体中文,繁体中文,英文,备注\n")

	//for i,content:= range c.content {
	//	for k,v:=range content {
	//		if i == 0 {
	//			_, _ = f.WriteString(k + "," + v+"\n")
	//		}else {
	//			//要再取出文件数据追加
	//
	//		}
	//	}
	//}

	var csvSlice []map[string][]string

	for _, content := range c.content {
		var s = make(map[string][]string)
		for k := range content {
			for i, j := range content {
				if k == i {
					s[k] = append(s[k], j)
				}
			}
		}
		csvSlice = append(csvSlice, s)
	}

	fmt.Println(csvSlice)
}
