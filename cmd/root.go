package cmd

import (
	"fmt"
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
	Short: "ki18n 是一个快速把excel[csv] 文件快速转成语言包的工具",
	Long:  "ki18n 是一个支持EXCEL，CSV 格式的语言文件快速转成JSON,PHP,IOS等平台的语言包工具",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("language:::", Languages)
	},
}

func init() {

	viper.SetConfigName("i18n")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	// 添加全局参数
	// 待转换的文件
	rootCmd.PersistentFlags().StringVarP(&Filename, "file", "f", "language.xlsx", "需要转换的源文件")
	// 需要转换的语言列表
	rootCmd.PersistentFlags().StringSliceVarP(&Languages, "language", "l", []string{"zh-CN", "zh-HK", "en-US"}, "转换目标的语言的列表")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

//获取 文件后缀
func Suffix(fullFilename string) string {
	filenameWithSuffix := path.Base(fullFilename) //获取文件名带后缀
	return path.Ext(filenameWithSuffix)           //获取文件后缀
}
