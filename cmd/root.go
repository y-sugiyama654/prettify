package cmd

import (
	"prettify/interface/cli"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "prettify",
	Short: "多機能フォーマッターCLIツール",
	Run: func(cmd *cobra.Command, args []string) {
		cli.Run(args)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
