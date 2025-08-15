package lookup

import (
	"encoding/json"
	"fmt"

	"github.com/Reggles44/secli/src/sec/company"
	"github.com/spf13/cobra"
)

var outputFields bool = false

var LookupCmd = &cobra.Command{
	Use:  "lookup [company or ticker] [tag]",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		company, err := company.Find(&args[0])
		if err != nil {
			panic(err)
		}

		if outputFields {
			for key := range company.Facts.USGaap {
				fmt.Println(key)
			}
			return
		}

		// fmt.Println(company.EntityName)

		facts := company.Facts.USGaap[args[1]]
		data, err := json.Marshal(facts)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(data))
	},
}

func init() {
	LookupCmd.Flags().BoolVar(&outputFields, "fields", false, "Outputs available fields")
}
