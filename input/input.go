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

//获取语言生成的语言列表
func Language(section string) []string {
	config, err := LoadConfigFile(ini)
	if err != nil {
		return language
	}

	val, _ := config.GetValue(section, "language")
	return strings.Split(val, ",")
}
