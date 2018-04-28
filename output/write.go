package output

import (
	"io/ioutil"
	"fmt"
	"os"
	"ki18n/input"
	"ki18n/driver"
)

// 输出语言包,合成一个语言包json
type Write struct {
	Driver driver.Driver
}

func (this *Write) PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//合并输出一个语言包
func (this *Write) all() {

	content := ToJsonMerge(this.Driver)
	err := ioutil.WriteFile("locales.json", content, 0755)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("语言包生成完成!")
}

// 输出语言包,每列一个包
func (this *Write) list(outType string) {
	langdir := "lang"
	if isExist, _ := this.PathExists(langdir); !isExist {
		err := os.Mkdir(langdir, 0755)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	for k, v := range input.Language() {
		content := To(outType, this.Driver.Parse(k))
		err := ioutil.WriteFile(langdir+"/"+v+"."+outType, content, 0755)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	fmt.Println("语言包生成完成!")
	return

}

// 输出
func (this *Write) Run(outType string) {
	if input.IsMerge() {
		this.all()
		return
	}
	this.list(outType)
}
