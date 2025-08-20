package lookup

import (
	"encoding/json"
	"fmt"

	"github.com/Reggles44/secli/internal/company"
	"github.com/spf13/cobra"
)

var outputFields bool = false

var LookupCmd = &cobra.Command{
	Use:  "lookup [search] [tag]",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		company, err := company.Find(args[0])
		if err != nil {
			panic(err)
		}

		// fmt.Println(company.CIK)

		submissions, err := company.LatestSubmission()
		if err != nil {
			panic(err)
		}
		

		filings := submissions.GetFilings("10-K")

		// fd, _ := json.Marshal(filings)
		// fmt.Println(string(fd))

		for _, filing := range filings {
			fstring, _ := json.Marshal(filing)
			fmt.Println(string(fstring))
		}
	},
}

func init() {
	LookupCmd.Flags().BoolVar(&outputFields, "", false, "Outputs available fields")
}
