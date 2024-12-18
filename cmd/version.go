package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version: v0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
