package typer

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type CSV struct {
	name string
}

// 返回实例
func NewCSV(filename string) CSV {
	return CSV{
		name: filename,
	}
}

//解析 csv
func (c CSV) Parse(col int) map[string]string {

	var lang = make(map[string]string)

	for k, v := range c.ReadFile() {
		if k == 0 {
			continue
		}

		if len(v) == 0 {
			continue
		}

		lang[v[0]] = v[col+1]
	}
	//fmt.Println(lang)
	return lang
}

// 读取文件返回二维数据
func (c CSV) ReadFile() (csvs [][]string) {
	csvs = make([][]string, 0)
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
			if len(rows) == 0 {
				return
			}
			return
		}

		csvs = append(csvs, rows)
	}
	return csvs
}
