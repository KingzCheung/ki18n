package driver

import (
	"encoding/csv"
	"fmt"
	. "github.com/KingzCheung/ki18n/input"
	"io"
	"os"
	"strings"
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
func (c *CSV) Parse(col int) map[string]string {

	var lang = make(map[string]string)

	for k, v := range c.ReadFile() {
		if k == 0 {
			continue
		}
		lang[v[0]] = v[col]
	}
	return lang
}

// 读取文件返回二维数据
func (c *CSV) ReadFile() (csvs [][]string) {

	file, err := os.Open(c.name)
	if err != nil {
		fmt.Println("读取 CSV 文件失败，错误信息：", err.Error())
	}
	defer file.Close()
	csvr := csv.NewReader(file)

	for {
		rows, err := csvr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return
		}

		for _, v := range rows {
			cols := strings.Split(v, Splitter(";"))
			csvs = append(csvs, cols)
		}
	}
	return csvs
}
