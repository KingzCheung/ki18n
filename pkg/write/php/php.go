package php

import (
	"fmt"
	"github.com/kingzcheung/ki18n/pkg/util"
	"os"
)

type Php struct {
	content string
}

func NewPhpFile() *Php {
	return new(Php)
}

// 生成 PHP 数组文件
func (p *Php) Write(path string) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = file.WriteString(p.content)
	if err != nil {
		fmt.Println(err)
		return
	}
	p.clear()
}

func (p *Php) clear() {
	p.content = ""
}

//生成PHP数组
func (p *Php) Format(fileType string, rows [][]string, index int) {
	p.content = "<?php\n"
	p.content += "return [\n"

	for r, row := range rows {
		if r == 0 {
			continue
		}
		for c := range row {
			if c != index+1 {
				continue
			}
			p.content += fmt.Sprintf("    \"%s\"=>\"%s\",\n", row[0], util.Escape(row[index+1]))

		}
	}

	p.content += "];\n"
}
