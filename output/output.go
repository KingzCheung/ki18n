package output

import (
	"encoding/json"
	"fmt"
	"strings"
)

var zy []string = []string{
	"\\",
	"\"",
}

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

func inArray(ver string, arr []string) (b bool) {

	for _, v := range arr {
		if ver == v {
			b = true
			break
		}
	}
	return
}

// 返回JSON 格式数据
func ToJson(lang interface{}) []byte {
	enc, err := json.Marshal(lang)
	if err != nil {
		fmt.Println(err.Error())
		return []byte(nil)
	}
	return enc
}

//返回 PHP 格式数据
func ToPHP(lang map[string]string) []byte {
	prefixString := "<?php\n return [\n"
	suffixString := "];"

	php := prefixString
	for k, v := range lang {
		php += "\"" + k + "\"=>\"" + format(v) + "\",\n"
	}
	php += suffixString

	return []byte(php)
}

func ToStrings(lang map[string]string) []byte {
	var str string
	for k, v := range lang {
		str += "\"" + k + "\" = \"" + format(v) + "\"\n"
	}
	return []byte(str)
}

// 返回[]BYTE数据
func To(outType string, lang map[string]string) []byte {
	switch strings.ToLower(outType) {
	case PHP:
		return ToPHP(lang)
	case IOS:
		return ToStrings(lang)
	default:
		return ToJson(lang)
	}
}
