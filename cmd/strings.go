// 生成 strings 格式的语言包，适用于iOS
// 生成的格式如下：
// ki18n strings -f=文件
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(stringsCmd)
}

var stringsCmd = &cobra.Command{
	Use:   "strings",
	Short: "生成 strings 格式的语言包",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("strings")
	},
}
