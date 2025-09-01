package taxonomy

type Taxonomy[T any] struct {
	DEI    DEI[T]    `json:"dei"`
	Invest Invest[T] `json:"invest"`
	SRT    SRT[T]    `json:"srt"`
	USGaap USGaap[T] `json:"us-gaap"`
}
