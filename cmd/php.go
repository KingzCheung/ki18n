// 生成 php 格式的语言包
// 格式如下：
// ki18n php -f 文件

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(phpCmd)
}

var phpCmd = &cobra.Command{
	Use:   "php",
	Short: "生成 php 格式的语言包",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("php")
	},
}
