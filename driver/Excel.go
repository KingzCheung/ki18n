package driver

import (
	"github.com/tealeg/xlsx"
	"fmt"
	"ki18n/input"
	"strings"
)

type Excel struct {
	Name string
}

func NewExcel(fileName string) *Excel {
	return &Excel{
		Name: fileName,
	}
}

func (this *Excel) ReadFile() *xlsx.File {
	xlFile, err := xlsx.OpenFile(this.Name)
	if err != nil {
		fmt.Println("文件加载错误:", err.Error())
	}

	return xlFile
}

//解析器
func (this *Excel) Parse(cel int) map[string]string {
	var lang = make(map[string]string, 5)

	for _, sheet := range this.ReadFile().Sheet {
		//遍历 行
		for rk, row := range sheet.Rows {
			if rk == 0 {
				continue
			}
			jsonKey := row.Cells[0].String()
			//去除空格
			jsonKey = strings.TrimRight(jsonKey, "\n")
			jsonKey = strings.TrimRight(jsonKey, "\r\n")
			jsonKey = strings.TrimRight(jsonKey, "\r")
			jsonKey = strings.Trim(jsonKey, " ")
			if len(row.Cells) < cel {
				lang[jsonKey] = ""
			}
			lang[jsonKey] = row.Cells[cel+1].String()
		}
	}

	return lang

}

//所有语言包合成一个JSON

func (this *Excel) ParseAll(section string) map[string]map[string]string {
	locales := make(map[string]map[string]string)

	for k, v := range input.Language(section) {
		locales[v] = this.Parse(k)
	}

	return locales
}
