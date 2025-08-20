package cmd

import (
	"fmt"
	"os"

	"github.com/Reggles44/secli/cmd/debug"
	"github.com/Reggles44/secli/cmd/lookup"
	"github.com/Reggles44/secli/cmd/preload"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "sec",
	Short: "A SEC CLI Tool",
}

var (
	company string
	email   string
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(0)
	}
}

func init() {
	// cobra.OnInitialize(config.InitConfig)

	// RootCmd Flags
	rootCmd.PersistentFlags().Bool("debug", false, "Enable debug")
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))

	// Add Sub Commands
	rootCmd.AddCommand(debug.DebugCmd)
	rootCmd.AddCommand(lookup.LookupCmd)
	rootCmd.AddCommand(preload.PreloadCmd)
}
