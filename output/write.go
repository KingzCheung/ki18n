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

func (this *Write) All(section string) {

	content := ToJsonMerge(section, this.Driver)
	err := ioutil.WriteFile("locales.json", content, 0755)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("语言包生成完成!")
}

// 输出语言包,每列一个包
func (this *Write) List(section string, outType string) {
	langdir := "lang"
	if isExist, _ := this.PathExists(langdir); !isExist {
		err := os.Mkdir(langdir, 0755)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	for k, v := range input.Language(section) {
		content := To(outType, this.Driver.Parse(k))
		err := ioutil.WriteFile(langdir+"/"+v+"."+outType, content, 0755)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	fmt.Println("语言包生成完成!")
	return

}
