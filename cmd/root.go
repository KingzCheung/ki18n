package cmd

import (
	"fmt"
	"github.com/KingzCheung/ki18n/pkg/output"
	"github.com/KingzCheung/ki18n/pkg/typer"
	"github.com/KingzCheung/ki18n/pkg/write"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path"
)

// 文件
var Filename string

// 语言包列表
var Languages []string

var rootCmd = &cobra.Command{
	Use:   "ki18n",
	Short: "ki18n 是一个快速把 EXCEL[CSV] 文件快速转成语言包的工具",
	Long:  "ki18n 是一个支持 EXCEL，CSV 格式的语言文件快速转成JSON,PHP,IOS等平台的语言包工具",
	Run: func(cmd *cobra.Command, args []string) {
		// 命令默认生成 JSON 语言包
		Run(func(col int, name string, o *output.Output) {
			o.ToJson(col).Write(name + ".json")
		})
	},
}

func init() {

	viper.SetConfigName("i18n")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
	//if err := viper.ReadInConfig(); err != nil {
	//	//fmt.Println(err)
	//}

	// 添加全局参数
	// 待转换的文件
	rootCmd.PersistentFlags().StringVarP(&Filename, "file", "f", "", "需要转换的源文件")
	// 需要转换的语言列表
	rootCmd.PersistentFlags().StringSliceVarP(&Languages, "language", "l", []string{}, "转换目标的语言的列表")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func getFileName() string {
	if Filename != "" {
		return Filename
	}

	if f := viper.GetString("file"); f != "" {
		return f
	}

	return "language.xlsx"
}

//获取 文件后缀
func Suffix(fullFilename string) string {
	filenameWithSuffix := path.Base(fullFilename) //获取文件名带后缀
	return path.Ext(filenameWithSuffix)           //获取文件后缀
}

//返回生成的语言列表，优先cli
func languages() []string {
	// 如果cli有参数，就用优先使用cli参数
	if 0 != len(Languages) {
		return Languages
	}
	// 如果cli没有参数，就取 i18n.yaml 的参数
	vl := viper.GetStringSlice("language")
	if 0 != len(vl) {
		return vl
	}
	// 如果 i18n.yaml 都没有定义，就取默认值
	return []string{"zh-CN", "zh-HK", "en-US"}

}

func Run(writeCallback func(col int, name string, o *output.Output)) {
	var oPut *output.Output

	// 判断文件是否存在
	exists, _ := write.PathExists(getFileName())
	if !exists {
		color.Danger.Println("错误: 文件 [\"" + getFileName() + "\"] 不存在！")
		return
	}

	switch Suffix(getFileName()) {
	case ".csv":
		oPut = output.New(typer.NewCSV(getFileName()))
	case ".xlsx":
		oPut = output.New(typer.NewExcel(getFileName()))
	}

	for k, v := range languages() {
		writeCallback(k, v, oPut)
		//oPut.ToJson(k).Write(v + suffName)
	}

	color.Cyan.Println(">>>>> 语言包生成成功啦 <<<<<")
}
