package main

import (
	"flag"
	"fmt"
	"ki18n/output"
	"strings"
	"errors"
	"ki18n/driver"
)

const (
	CSV  = "csv"
	XLSX = "xlsx"
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

	d, err := FileType(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	write := output.Write{
		d,
	}

	if merge {
		write.All(configFile)
	} else {
		write.List(configFile, outputType)
	}

}

func FileType(filename string) (driver.Driver, error) {
	strArr := strings.Split(filename, ".")
	last := len(strArr) - 1
	if last < 0 {
		return nil, errors.New("文件名后缀错误")
	}
	var d driver.Driver
	switch strings.ToLower(strArr[last]) {
	case CSV:
		d = driver.NewCSV(filename)
	case XLSX:
		d = driver.NewExcel(filename)
	default:
		d = driver.NewExcel(filename)
	}
	return d, nil
}
