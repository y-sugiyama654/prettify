package cmd

import (
	"prettify/interface/cli"

	"github.com/spf13/cobra"
)

var formatType string

var rootCmd = &cobra.Command{
	Use:   "prettify",
	Short: "多機能フォーマッターCLIツール(JSON / SQL)",
	Run: func(cmd *cobra.Command, args []string) {
		cli.Run(formatType, args)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&formatType, "type", "t", "json", "整形対象のタイプ（json / sql）")
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
