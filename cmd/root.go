package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "ki18n",
	Short: "ki18n is a parse excel to i18n files",
	Long:  `目标是解析一定格式的语言文件，来生成对应项目中可直接使用的语言包文件`,
	Run: func(cmd *cobra.Command, args []string) {
		//添加东西到这里来
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
