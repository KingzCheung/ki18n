package cmd

import (
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

const version = "2.0.0 beta"

var (
	versionCmd = &cobra.Command{

		Use:   "version",
		Short: "Print version information and quit",
		Run: func(cmd *cobra.Command, args []string) {
			color.Green.Printf("ki18n version:%s\n", version)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
