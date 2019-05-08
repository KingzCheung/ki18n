package i18n

import (
	"errors"
	"fmt"
	"github.com/KingzCheung/ki18n/pkg/read/csv"
	"github.com/KingzCheung/ki18n/pkg/read/excel"
	"github.com/KingzCheung/ki18n/pkg/util"
	"github.com/KingzCheung/ki18n/pkg/write/json"
	"github.com/KingzCheung/ki18n/pkg/write/php"
	strings2 "github.com/KingzCheung/ki18n/pkg/write/strings"
	"github.com/gookit/color"
	"os"
	"strings"
)

const (
	JSON    = "json"
	PHP     = "php"
	STRINGS = "strings"
	DistDir = "lang"
	iOS     = "ios"
	OBJC    = "objc"
	Swift   = "swift"
)

// 步骤

//读取源文件
//解析文件内容
//生成对应的数组文件
type Reader interface {
	Read() ([][]string, error)
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
		fileType: fileType(format),
	}
}

// fileType 适配不同的strings命令格式别名
func fileType(t string) string {
	switch t {
	case STRINGS, iOS, Swift, OBJC:
		return STRINGS
	default:
		return t
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
		return nil, errors.New("unrecognized table file")
	}
}

// writeInstance 获取写实例,默认获取json的实例
// format 传入需要生成的格式
func writeInstance(format string) Writer {
	switch format {
	case JSON:
		return json.NewJsonFile()
	case PHP:
		return php.NewPhpFile()
	case STRINGS, iOS, Swift, OBJC:
		return strings2.NewStringsFile()
	default:
		return json.NewJsonFile()
	}
}

func (i *I18n) ParseToFile() {

	//获取读取的数组
	rows, err := i.read.Read()

	if err != nil {
		color.Red.Println(err)
		return
	}

	if len(rows) == 0 {
		color.Red.Println("No data is read, please check the Sheet name")
		return
	}
	//生成目录
	if b, _ := util.PathExists(DistDir); !b {
		err = os.Mkdir(DistDir, 0755)
		if err != nil {
			color.Red.Println(err)
			return
		}
	}
	languages := i.lang
	//如果语言列表为空，就取表头
	if len(languages) == 0 {
		languages = rows[0][1:]
	}
	for k, v := range languages {
		i.write.Format(i.fileType, rows, k)
		wFileName := fmt.Sprintf("%s/%s.%s", DistDir, v, i.fileType)
		i.write.Write(wFileName)
	}

}
