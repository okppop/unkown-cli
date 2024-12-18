package cmd

import (
	"cli/handler"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init database",
	Long:  "init will create sqlite database file",
	Run:   handler.InitCmd,
}

func init() {
	rootCmd.AddCommand(initCmd)
}
