package strings

import (
	"fmt"
	"github.com/KingzCheung/ki18n/pkg/util"
	"os"
)

type Strings struct {
	content string
}

func (s *Strings) Write(path string) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = file.WriteString(s.content)
	if err != nil {
		fmt.Println(err)
		return
	}
	s.clear()
}

func (s *Strings) clear() {
	s.content = ""
}

func (s *Strings) Format(fileType string, rows [][]string, index int) {
	for r, row := range rows {
		if r == 0 {
			continue
		}
		for c := range row {
			if c != index+1 {
				continue
			}
			s.content += fmt.Sprintf("\"%s\"=\"%s\";\n", row[0], util.Escape(row[index+1]))
		}
	}
}

func NewStringsFile() *Strings {
	return new(Strings)
}
