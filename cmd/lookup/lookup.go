package lookup

import (
	"encoding/json"
	"fmt"

	"github.com/Reggles44/secli/src/sec/company"
	"github.com/spf13/cobra"
)

var LookupCmd = &cobra.Command{
	Use:  "lookup [company or ticker] [form]",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		company, err := company.Find(&args[0])
		if err != nil {
			panic(err)
		}

		data, err := json.Marshal(company)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(data))
	},
}
