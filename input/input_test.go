package input

import (
	"testing"
	"strings"
)

func TestLanguage(t *testing.T)  {
	languageArr := Language()
	lang := "zh-CN,zh-HK,en-US"
	if strings.Join(languageArr, ",") != lang {
		t.Error("获取语言包列表有误")
	}
}