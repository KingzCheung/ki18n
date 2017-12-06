package main

import (
	"flag"
	"fmt"
	output2 "ki18n/output"
	"ki18n/driver"
)

func main() {
	//获取文件名
	var filename string
	var version bool
	var merge bool
	var configFile string
	var outputType string

	flag.StringVar(&filename, "f", "language.xlsx", "解析的文件名")
	flag.StringVar(&configFile, "c", "default", "i18n.ini 的配置选项")
	flag.StringVar(&outputType, "t", "json", "生成文件类型,支持[json,php,strings]")
	flag.BoolVar(&merge, "m", false, "是否合并语言包")
	flag.BoolVar(&version, "v", false, "版本")
	flag.Parse()

	if version {
		fmt.Println("version:0.01")
		return
	}

	output := output2.NewOutput(filename)

	if merge {
		output.All(configFile)
	} else {
		output.List(configFile, outputType)
	}

}


