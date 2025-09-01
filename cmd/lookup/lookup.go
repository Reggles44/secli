package lookup

import (
	"github.com/spf13/cobra"
)

var outputFields bool = false

var LookupCmd = &cobra.Command{
	Use:  "lookup [search] [tag]",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// companySearchString := args[0]
		// filingStringSearch := args[1]
		//
		// company, err := company.Find(companySearchString)
		// if err != nil {
		// 	panic(err)
		// }
		//
		// fmt.Println(company.CIK)
		//
		// submissions, err := company.LatestSubmission()
		// if err != nil {
		// 	panic(err)
		// }
	},
}

func init() {
	LookupCmd.Flags().BoolVar(&outputFields, "", false, "Outputs available fields")
}
