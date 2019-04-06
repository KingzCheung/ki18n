package php

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gookit/color"
	"github.com/z7zmey/php-parser/php7"
	"github.com/z7zmey/php-parser/visitor"
	"io/ioutil"
	"os"
)

type Php struct {
	dir string
}

func NewPhp(dir string) *Php {
	return &Php{dir: dir}
}

//DirRead 读取文件下的内容
func (p *Php) DirRead() []map[string]string {
	//读取 DIR 目录下的语言包文件
	dirList, err := ioutil.ReadDir(p.dir)

	if err != nil {
		color.Red.Println(err)
		return []map[string]string{}
	}
	languages := make([]map[string]string, 0)

	for _, d := range dirList {
		file, _ := os.Open(p.dir + "/" + d.Name())
		phpStr, _ := ioutil.ReadAll(file)
		src := bytes.NewBufferString(string(phpStr))

		php7parser := php7.NewParser(src, "test.php")
		php7parser.WithFreeFloating()
		php7parser.Parse()
		for _, e := range php7parser.GetErrors() {
			fmt.Println(e)
		}
		nodes := php7parser.GetRootNode()
		nsResolver := visitor.NewNamespaceResolver()
		nodes.Walk(nsResolver)

		buf := new(bytes.Buffer)

		dumper := &visitor.PrettyJsonDumper{
			Writer: buf,
		}

		nodes.Walk(dumper)

		var ast Ast
		err = json.Unmarshal(buf.Bytes(), &ast)
		if err != nil {
			color.Red.Println(err)
			return []map[string]string{}
		}

		lang := make(map[string]string)

		for _, item := range ast.Stmts[0].Expr.Items {
			if item.Key.Value == "" {
				continue
			}
			lang[item.Key.Value] = item.Val.Value
		}

		languages = append(languages, lang)

	}

	fmt.Println(languages)

	return languages
}
