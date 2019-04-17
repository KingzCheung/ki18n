package csv

import (
	"encoding/csv"
	"github.com/gookit/color"
	"io"
	"os"
)

type Csv struct {
	filePath string
}

func NewCsv(filePath string) *Csv {
	return &Csv{filePath: filePath}
}

func (c *Csv) Read() [][]string {
	file, err := os.Open(c.filePath)
	if err != nil {
		color.Red.Println(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	csvRes := make([][]string, 0)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			color.Red.Println(err)
			return [][]string{}
		}
		csvRes = append(csvRes, record)
	}
	return csvRes
}
