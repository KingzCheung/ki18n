package php

import (
	"fmt"
	"github.com/KingzCheung/ki18n/pkg/util"
	"os"
)

type PhpFile struct {
	content string
}

func NewPhpFile() *PhpFile {
	return new(PhpFile)
}

// 生成 PHP 数组文件
func (p *PhpFile) Write(path string) {
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

func (p *PhpFile) clear() {
	p.content = ""
}

//生成PHP数组
func (p *PhpFile) Format(fileType string, rows [][]string, index int) {
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
