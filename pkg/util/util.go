package util

import (
	"os"
	"strings"
)

//转义内容
func Escape(source string) (dist string) {
	//转义双引号
	dist = strings.Replace(source, "\"", "\\\"", -1)
	return
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// InSlice 判断 一个值是否存在于一个slice中
func InSlice(slices []string, search string) bool {
	for _, s := range slices {
		if s == search {
			return true
		}
	}
	return false
}
