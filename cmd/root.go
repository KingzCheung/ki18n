package cmd

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"os"
)

var v bool
var rootCmd = &cobra.Command{
	Use:   "ki18n",
	Short: "ki18n is a parse excel to i18n files",
	Long:  `目标是解析一定格式的语言文件，来生成对应项目中可直接使用的语言包文件`,
	Run: func(cmd *cobra.Command, args []string) {
		//添加东西到这里来

		if v {
			color.Green.Printf("version: %s", version)
		}
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&v, "version", "v", false, "Print version information and quit")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
