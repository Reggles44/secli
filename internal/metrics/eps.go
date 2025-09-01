package metrics

import (
	"fmt"

	"github.com/Reggles44/secli/internal/docs"
)

func EPS(d *docs.Doc) float64 {
	netIncome := d.Taxonomy.USGaap.NetIncomeLoss.Value.Value
	dividends := d.Taxonomy.USGaap.OtherPreferredStockDividendsAndAdjustments.Value.Value
	shares := d.Taxonomy.DEI.EntityCommonStockSharesOutstanding.Value.Value

	fmt.Println(d.AccessionNumber, d.FilingDate, d.Form)
	fmt.Printf("Net Income %v\n", netIncome)
	fmt.Printf("Dividends %v\n", dividends)
	fmt.Printf("Shares %v\n", shares)
	return (netIncome - dividends) / shares

}
