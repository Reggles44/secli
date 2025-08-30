package taxonomy

type DEI[T any] struct {
	EntityCommonStockSharesOutstanding T `json:"EntityCommonStockSharesOutstanding"`
	EntityPublicFloat                  T `json:"EntityPublicFloat"`
}
