package write

import (
	"fmt"
	"io/ioutil"
	"os"
)

// 判断路径是否存在
func pathExists(path string) (bool, error) {
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
	langDir := "lang"
	if isExist, _ := pathExists(langDir); !isExist {
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
