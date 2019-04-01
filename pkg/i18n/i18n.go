package i18n

import (
	"fmt"
	"github.com/KingzCheung/ki18n/pkg/read/excel"
	"github.com/KingzCheung/ki18n/pkg/util"
	"github.com/KingzCheung/ki18n/pkg/write/json"
	"github.com/KingzCheung/ki18n/pkg/write/php"
	"os"
)

const (
	JSON     = "json"
	PHP      = "php"
	STRINGS  = "strings"
	DIST_DIR = "lang"
)

// 步骤

//读取源文件
//解析文件内容
//生成对应的数组文件
type Reader interface {
	Read() [][]string
}

type Writer interface {
	Write(path string)
	Format(fileType string, c [][]string, index int)
}

type I18n struct {
	read     Reader
	write    Writer
	fileType string
	lang     []string
}

func NewI18n(srcPath string, lang []string, format string) *I18n {
	read := excel.NewExcel(srcPath)
	write := getWriteInstance(format)
	return &I18n{
		read:     read,
		write:    write,
		lang:     lang,
		fileType: format,
	}
}

func getWriteInstance(format string) Writer {
	switch format {
	case JSON:
		return json.NewJsonFile()
	case PHP:
		return php.NewPhpFile()
	default:
		return json.NewJsonFile()
	}
}

func (i *I18n) ParseToFile() {
	//生成目录
	if b, _ := util.PathExists(DIST_DIR); !b {
		_ = os.Mkdir(DIST_DIR, 0755)
	}
	//获取读取的数组
	rows := i.read.Read()

	for k, v := range i.lang {
		i.write.Format(i.fileType, rows, k)
		wfileName := fmt.Sprintf("lang/%s.%s", v, i.fileType)
		i.write.Write(wfileName)
	}

}
