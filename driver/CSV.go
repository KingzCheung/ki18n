package driver

import (
	"os"
	"fmt"
	"encoding/csv"
	"io"
	"strings"
	"ki18n/input"
)

type CSV struct {
	name string
}

// 返回实例
func NewCSV(filename string) *CSV {
	return &CSV{
		name: filename,
	}
}

//解析 csv
func (this *CSV) Parse(col int) map[string]string {

	var lang = make(map[string]string)

	for _, v := range this.ReadFile() {
		lang[v[0]] = v[col]
	}
	return lang
}

// 读取文件返回二维数据
func (this *CSV) ReadFile() (csvs [][]string) {

	file, err := os.Open(this.name)
	if err != nil {
		fmt.Println("读取 CSV 文件失败，错误信息：", err.Error())
	}
	defer file.Close()
	csvr := csv.NewReader(file)

	for {
		rows, err := csvr.Read()
		if err == io.EOF {
			break
		} else if (err != nil) {
			return
		}

		for _, v := range rows {
			cols := strings.Split(v, input.Splitter())
			csvs = append(csvs, cols)
		}
	}
	return csvs
}
