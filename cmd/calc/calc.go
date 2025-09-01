package calc

import (
	"fmt"

	"github.com/Reggles44/secli/internal/company"
	"github.com/Reggles44/secli/internal/forms"
	"github.com/Reggles44/secli/internal/metrics"
	"github.com/spf13/cobra"
)

var CalcCmd = &cobra.Command{
	Use:  "calc [search] ",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// method := args[0]
		compSearch := args[0]

		c, err := company.Find(compSearch)
		if err != nil {
			panic(err)
		}

		docs, err := c.Docs()
		if err != nil {
			panic(err)
		}

		latestTenK := docs[forms.TenQ][0]
		if err := latestTenK.Fill(); err != nil {
			panic(err)
		}

		// fmt.Println(latestTenK.AccessionNumber)
		// fmt.Printf("Net Income Loss %+v\n", latestTenK.Taxonomy.USGaap.NetIncomeLoss)

		eps, err := metrics.EPS(latestTenK)
		if err != nil {
			panic(err)
		}
		fmt.Println(eps)
	},
}
