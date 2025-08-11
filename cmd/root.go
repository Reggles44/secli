package cmd

import (
	"fmt"
	"os"

	"github.com/Reggles44/secli/cmd/lookup"
	"github.com/Reggles44/secli/src/config"
	"github.com/spf13/cobra"
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
	cobra.OnInitialize(config.InitConfig)

	// RootCmd Flags
	rootCmd.PersistentFlags().Bool("test", true, "Set testing mode")
	// rootCmd.PersistentFlags().StringVar("test", true, "Set testing mode")

	// Add Sub Commands
	rootCmd.AddCommand(lookup.LookupCmd)
}
