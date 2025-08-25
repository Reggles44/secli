package calc

import (
	"fmt"

	"github.com/Reggles44/secli/internal/company"
	"github.com/Reggles44/secli/internal/company/metrics"
	"github.com/spf13/cobra"
)

var CalcCmd = &cobra.Command{
	Use:  "calc [search] ",
	Args: cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		// method := args[0]
		compSearch := args[1]
		form := args[2]

		c, err := company.Find(compSearch)
		if err != nil {
			panic(err)
		}

		result, err := metrics.EBITDA(c, form)
		if err != nil {
			panic(err)
		}

		fmt.Println(result)
	},
}
