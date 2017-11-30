package output

import (
	"encoding/json"
	"fmt"
	"strings"
)

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
		php += "'" + k + "'=>'" + v + "',\n"
	}
	php += suffixString

	return []byte(php)
}

func ToStrings(lang map[string]string) []byte {
	var str string
	for k, v := range lang {
		str += "\"" + k + "\" = \"" + v + "\"\n"
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
