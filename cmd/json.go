// 合成 JSON 格式的子命令
// 用法如下:
// ki18n json [-f lang.xlsx -m]

package cmd

import (
	"fmt"
	"github.com/KingzCheung/ki18n/pkg/output"
	"github.com/KingzCheung/ki18n/pkg/typer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		//fmt.Println("json")

	},
}

func languages() []string {
	if Languages == nil {
		return viper.GetStringSlice("language")
	}
	return Languages
}

func Run() {
	var oput *output.Output
	switch Suffix(Filename) {
	case "csv":
		oput = output.New(typer.NewCSV(Filename))
	case "xlsx":
		oput = output.New(typer.NewExcel(Filename))
	}

	for k, v := range languages() {
		fmt.Println(k, v)
		oput.ToJson(k).Write(v)
	}

}
