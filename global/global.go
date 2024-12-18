package global

import (
	"fmt"
	"os"
)

var (
	DataPath string
)

func init() {
	homePath, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	DataPath = homePath + "/.local/share" + "/pmtest"
}
