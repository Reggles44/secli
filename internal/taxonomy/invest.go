package taxonomy

type Invest[T any] struct {
	DerivativeNotionalAmount        T `json:"DerivativeNotionalAmount"`
	InvestmentWarrantsExercisePrice T `json:"InvestmentWarrantsExercisePrice"`
}
