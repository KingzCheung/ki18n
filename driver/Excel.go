package driver

import (
	"github.com/tealeg/xlsx"
	"fmt"
	"encoding/json"
	"os"
	"io/ioutil"
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

func (this *Excel) PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//解析器
func (this *Excel) parse(cel int) map[string]string {
	var lang = make(map[string]string, 5)

	for _, sheet := range this.ReadFile().Sheet {
		//遍历 行
		for rk, row := range sheet.Rows {
			if rk == 0 {
				continue
			}
			jsonKey := row.Cells[0].String()
			//去除空格
			jsonKey = strings.Trim(jsonKey,"\n")
			jsonKey = strings.Trim(jsonKey,"\r\n")
			jsonKey = strings.Trim(jsonKey,"\r")
			jsonKey = strings.Trim(jsonKey," ")
			if len(row.Cells) < cel {
				lang[jsonKey] = ""
			}
			lang[jsonKey] = row.Cells[cel+1].String()
		}
	}

	return lang

}

//所有语言包合成一个JSON

func (this *Excel) parseAll(section string) map[string]map[string]string {
	locales := make(map[string]map[string]string)

	for k, v := range input.Language(section) {
		locales[v] = this.parse(k)
	}

	return locales
}

// 转换成JSON
func (this *Excel) toJson(lang interface{}) []byte {
	enc, err := json.Marshal(lang)
	if err != nil {
		fmt.Println(err.Error())
		return []byte(nil)
	}

	return enc
}

// 输出语言包,合成一个语言包json

func (this *Excel) OutputAll(section string) {
	content := this.toJson(this.parseAll(section))
	err := ioutil.WriteFile("locales.json", content, 0755)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	
	fmt.Println("语言包生成完成!")
}

// 输出语言包,每列一个包
func (this *Excel) Output(section string) {
	langdir := "lang"
	if isExist, _ := this.PathExists(langdir); !isExist {
		err := os.Mkdir(langdir, 0755)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	for k, v := range input.Language(section) {
		content := this.toJson(this.parse(k))
		err := ioutil.WriteFile(langdir+"/"+v+".json", content, 0755)
		if err != nil {
			fmt.Println(err.Error())
		}

	}

	fmt.Println("语言包生成完成!")
	return

}
