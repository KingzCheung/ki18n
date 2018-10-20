// 从输入流中处理数据，并通过不同的参数输出不同的数据格式

package output

import (
	"encoding/json"
	"github.com/KingzCheung/ki18n/pkg/typer"
	"github.com/KingzCheung/ki18n/pkg/write"
	"log"
)

var zy = []string{
	"\\",
	"\"",
}

type Output struct {
	data          map[string]string
	output        []byte
	inputFileType typer.Typer
}

func New(t typer.Typer) *Output {
	return &Output{
		inputFileType: t,
	}
}

// 把数据写到文件
func (o *Output) Write(name string) {
	write.Write(name, o.output)
}

// 返回JSON 格式数据
func (o *Output) ToJson(col int) *Output {
	var err error
	// 解析出数组
	o.data = o.inputFileType.Parse(col)
	o.output, err = json.Marshal(o.data)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return o
}

//返回 PHP 格式数据
func (o *Output) ToPHP() *Output {
	prefixString := "<?php\n return [\n"
	suffixString := "];"

	php := prefixString
	for k, v := range o.data {
		php += "\"" + k + "\"=>\"" + format(v) + "\",\n"

	}
	php += suffixString
	o.output = []byte(php)
	return o
}

func (o *Output) ToStrings(lang map[string]string) *Output {
	var str string
	for k, v := range o.data {
		str += "\"" + k + "\" = \"" + format(v) + "\"\n"
	}
	o.output = []byte(str)
	return o
}

// 判断数组是否在某一个字符

func inArray(ver string, arr []string) (b bool) {

	for _, v := range arr {
		if ver == v {
			b = true
			break
		}
	}
	return
}

// 格式化数据
func format(str string) string {
	var newStr string
	for _, v := range str {
		if inArray(string(v), zy) {
			newStr += "\\" + string(v)
		} else {
			newStr += string(v)
		}
	}
	return newStr
}
