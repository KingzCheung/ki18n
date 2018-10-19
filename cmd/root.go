package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var filename string

var rootCmd = &cobra.Command{
	Use:   "ki18n",
	Short: "ki18n 是一个快速把excel[csv] 文件快速转成语言包的工具",
	Long:  "ki18n 支持EXCEL，CSV 格式的语言文件快速黑的成JSON,PHP,IOS等平台的语言包工具",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ki18n")
		fmt.Println(filename)
	},
}

func init() {
	//添加全局参数
	rootCmd.PersistentFlags().StringVarP(&filename, "file", "f", "language.xlsx", "需要转换的源文件")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
