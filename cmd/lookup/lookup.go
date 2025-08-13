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
		submission, err := company.Find(&args[0])
		if err != nil {
			panic(err)
		}

		if len(args) == 1 {
			data, err := json.Marshal(submission)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(data))
		}

		data, err := submission.GetFile(args[1])
		if err != nil {
			panic(err)
		}
		fmt.Println(string(*data))
	},
}
