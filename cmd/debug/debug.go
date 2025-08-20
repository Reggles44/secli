package debug

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var DebugCmd = &cobra.Command{
	Use:  "debug",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(viper.GetBool("debug"))
	},
}


