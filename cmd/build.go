package cmd

import (
	i18n2 "github.com/KingzCheung/ki18n/pkg/i18n"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "构建项目可使用的i18n文件",
	Long:  `可通过 xlsx 或者 csv 格式的管理的语言包生成目标项目可使用的i18n文件[json,php]`,
	Run: func(cmd *cobra.Command, args []string) {
		//运行build指令
		i18n := i18n2.NewI18n(args[0], lang, format)
		i18n.ParseToFile()
	},
}

// 生成的格式
var format string

// 生成的语言
var lang []string

func init() {
	buildCmd.Flags().StringSliceVarP(&lang, "lang", "l", []string{"zh-cn"}, "生成的语言包列表")
	buildCmd.Flags().StringVarP(&format, "format", "", "json", "生成的格式")
	rootCmd.AddCommand(buildCmd)
}
