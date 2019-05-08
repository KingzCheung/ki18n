package excel

import (
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
		workSheet: "Sheet1",
	}
}

//读取excel文件

func (e *Excel) Read() ([][]string, error) {
	xlsx, err := excelize.OpenFile(e.filePath)
	if err != nil {
		return [][]string{}, err
	}
	return xlsx.GetRows(e.workSheet), nil
}
