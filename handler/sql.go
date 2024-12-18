package handler

import (
	"cli/global"
	"fmt"

	"github.com/spf13/cobra"
)

func InitCmd(cmd *cobra.Command, args []string) {
	fmt.Printf("AppDataPath : %s", global.DataPath)
	return
}
