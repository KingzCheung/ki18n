package i18n

import (
	"errors"
	"fmt"
	"github.com/KingzCheung/ki18n/pkg/read/csv"
	"github.com/KingzCheung/ki18n/pkg/read/excel"
	"github.com/KingzCheung/ki18n/pkg/util"
	"github.com/KingzCheung/ki18n/pkg/write/json"
	"github.com/KingzCheung/ki18n/pkg/write/php"
	"github.com/gookit/color"
	"os"
	"strings"
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
	read, err := readInstance(srcPath)
	if err != nil {
		color.Red.Println(err)
		return &I18n{}
	}
	write := writeInstance(format)
	return &I18n{
		read:     read,
		write:    write,
		lang:     lang,
		fileType: format,
	}
}

//获取读实例
func readInstance(srcPath string) (Reader, error) {
	s := strings.Split(srcPath, ".")
	suffix := s[len(s)-1]
	switch suffix {
	case "xlsx":
		return excel.NewExcel(srcPath), nil
	case "csv":
		return csv.NewCsv(srcPath), nil
	case "xls":
		return excel.NewExcel(srcPath), nil
	default:
		return nil, errors.New("无法识别表格文件")
	}
}

//获取写实例
func writeInstance(format string) Writer {
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
	languages := i.lang
	//如果语言列表为空，就取表头
	if len(languages) == 0 {
		languages = rows[0][1:]
	}

	for k, v := range languages {
		i.write.Format(i.fileType, rows, k)
		wFileName := fmt.Sprintf("lang/%s.%s", v, i.fileType)
		i.write.Write(wFileName)
	}

}
