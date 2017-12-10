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

func TestSplitter(t *testing.T) {
	splitter := ";"
	if splitter != Splitter(";") {
		t.Error("配置解析错误")
	}
}

func TestIsMerge(t *testing.T) {
	if IsMerge() != false {
		 t.Error("配置解析错误")
	}
}