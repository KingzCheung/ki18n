// 合成 JSON 格式的子命令
// 用法如下:
// ki18n json [-f lang.xlsx -m]

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// 是否要合并整个语言包为一个JSON文件
var merge bool

func init() {
	rootCmd.AddCommand(jsonCmd)
	jsonCmd.Flags().BoolVarP(&merge, "merge", "m", false, "是否合并")
}

var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "生成 json 格式的语言包",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("json")
	},
}
