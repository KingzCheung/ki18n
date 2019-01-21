package typer

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"strings"
)

type Excel struct {
	Name string
}

func NewExcel(fileName string) Excel {
	return Excel{
		Name: fileName,
	}
}

func (e Excel) ReadFile() *xlsx.File {

	xlFile, err := xlsx.OpenFile(e.Name)
	if err != nil {
		fmt.Println("文件加载错误")
		return nil
	}
	return xlFile
}

//解析器
func (e Excel) Parse(col int) map[string]string {
	var lang = make(map[string]string, 5)
	var jsonKey string
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	for _, sheet := range e.ReadFile().Sheet {
		//遍历 行
		for rk, row := range sheet.Rows {
			if rk == 0 {
				continue
			}
			if len(row.Cells) > 0 {
				jsonKey = row.Cells[0].String()
			} else {
				continue
			}
			//去除空格
			jsonKey = strings.TrimRight(jsonKey, "\n")
			jsonKey = strings.TrimRight(jsonKey, "\r\n")
			jsonKey = strings.TrimRight(jsonKey, "\r")
			jsonKey = strings.Trim(jsonKey, " ")
			if strings.Trim(jsonKey, " ") == "" {
				continue
			}
			if len(row.Cells) == 0 {
				lang[jsonKey] = ""
			} else {
				lang[jsonKey] = row.Cells[col+1].String()

			}
		}
	}

	return lang

}
