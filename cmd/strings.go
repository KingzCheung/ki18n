// 生成 strings 格式的语言包，适用于iOS
// 生成的格式如下：
// ki18n strings -f=文件
package cmd

import (
	"github.com/KingzCheung/ki18n/pkg/output"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(stringsCmd)
}

var stringsCmd = &cobra.Command{
	Use:     "strings",
	Aliases: []string{"ios", "swift", "objc"},
	Short:   "生成 strings 格式的语言包",
	Run: func(cmd *cobra.Command, args []string) {
		Run(func(col int, name string, o *output.Output) {
			o.ToStrings(col).WriteWithDir(name)
		})
	},
}
