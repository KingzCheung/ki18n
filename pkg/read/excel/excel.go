package excel

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

type Excel struct {
	filePath  string
	workSheet string
}

//工作表1实例
func NewExcel(filePath string) *Excel {
	return &Excel{
		filePath:  filePath,
		workSheet: "工作表1",
	}
}

//读取excel文件

func (e *Excel) Read() [][]string {
	xlsx, err := excelize.OpenFile(e.filePath)
	if err != nil {
		fmt.Println(err)
		return [][]string{}
	}
	return xlsx.GetRows(e.workSheet)
}
