package excel

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gookit/color"
)

type Excel struct {
	filePath  string
	workSheet string
}

//工作表1实例
func NewExcel(filePath string) *Excel {
	return &Excel{
		filePath:  filePath,
		workSheet: "Sheet1",
	}
}

//读取excel文件

func (e *Excel) Read() [][]string {
	xlsx, err := excelize.OpenFile(e.filePath)
	if err != nil {
		color.Red.Println(err)
		return [][]string{}
	}
	return xlsx.GetRows(e.workSheet)
}
