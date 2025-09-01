package metrics

import (
	"fmt"

	"github.com/Reggles44/secli/internal/docs"
)

func EPS(d *docs.Doc) (float64, error) {
	netIncome, err := d.Taxonomy.USGaap.NetIncomeLoss.YearlyValue()
	if err != nil {
		return 0, err
	}

	dividends, err:= d.Taxonomy.USGaap.OtherPreferredStockDividendsAndAdjustments.Value()
	if err != nil {
		dividends = 0
	}

	shares, err := d.Taxonomy.DEI.EntityCommonStockSharesOutstanding.Value()
	if err != nil {
		return 0, err
	}


	fmt.Println(d.AccessionNumber, d.FilingDate, d.Form)
	fmt.Printf("Net Income %v\n", netIncome)
	fmt.Printf("Dividends %v\n", dividends)
	fmt.Printf("Shares %v\n", shares)
	return (netIncome - dividends) / shares, nil
}
