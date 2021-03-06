package json

import (
	"fmt"
	"github.com/KingzCheung/ki18n/pkg/util"
	"os"
	"strings"
)

type Json struct {
	content string
}

func NewJsonFile() *Json {
	return new(Json)
}

//写入文件
func (j *Json) Write(path string) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = file.WriteString(j.content)
	if err != nil {
		fmt.Println(err)
	}
	j.clear()
}

// 清空内容
func (j *Json) clear() {
	j.content = ""
}

//格式化数据
func (j *Json) Format(fileType string, rows [][]string, index int) {
	j.content = "{"

	for r, row := range rows {
		if r == 0 {
			continue
		}
		for c := range row {
			if c != index+1 {
				continue
			}
			j.content += fmt.Sprintf("\"%s\":\"%s\",", row[0], util.Escape(row[index+1]))
			if r == len(rows)-1 {
				j.content = strings.TrimRight(j.content, ",")
			}
		}
	}

	j.content += "}"
}
