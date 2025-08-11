package lookup

import (
	"github.com/Reggles44/secli/src/sec/company"
	"github.com/spf13/cobra"
)

var LookupCmd = &cobra.Command{
	Use:  "lookup [company or ticker]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		company.Find(&args[0])
	},
}
