package lookup

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Reggles44/secli/src/sec/company"
	"github.com/spf13/cobra"
)

var lookupCmd = &cobra.Command{
	Use: "lookup",
	Run: func(cmd *cobra.Command, args []string) {
		company, err := company.Find(args[0])
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(0)
		}

		jsonData, err := json.Marshal(company)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(0)
		}

		fmt.Fprint(os.Stdout, string(jsonData))
	},
}
