package input

import (
	. "github.com/Unknwon/goconfig"
	"strings"
)

var ini string = "i18n.ini"
var language []string = []string{
	"zh-CN",
	"zh-HK",
	"en-US",
}
var section string = "DEFAULT"

//获取语言生成的语言列表
func Language() []string {
	config, err := LoadConfigFile(ini)
	if err != nil {
		return language
	}

	val, _ := config.GetValue(section, "language")
	return strings.Split(val, ",")
}

func Splitter() string {
	config, err := LoadConfigFile(ini)
	if err != nil {
		return ";"
	}
	val, _ := config.GetValue(section, "splitter")
	return val
}

func IsMerge() bool {
	config, err := LoadConfigFile(ini)
	if err != nil {
		return false
	}
	return config.MustBool(section, "merge", false)
}
