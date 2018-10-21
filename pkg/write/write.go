package write

import (
	"fmt"
	"io/ioutil"
	"os"
)

var langDir = "lang"

// 判断路径是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 把数据写入文件
func Write(name string, content []byte) {
	if isExist, _ := PathExists(langDir); !isExist {
		err := os.Mkdir(langDir, 0755)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	err := ioutil.WriteFile(langDir+"/"+name, content, 0755)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

// 写入包涵文件夹格式的语言包
func WriteWithDir(name string, content []byte) {

	langFileName := "Localizable"

	if isExist, _ := PathExists(langDir); !isExist {
		err := os.Mkdir(langDir, 0755)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	langDirName := langDir + "/" + name + ".lproj"
	if isExist, _ := PathExists(langDirName); !isExist {
		err := os.Mkdir(langDirName, 0755)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	//f := langDirName + "/" + langFileName
	err := ioutil.WriteFile(langDirName+"/"+langFileName+".strings", content, 0755)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
