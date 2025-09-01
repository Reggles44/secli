package facts

type FactSearch struct {
	Taxonomy string
	Field    string
	Form     string
	TTM      bool
}

func (f *Facts) Find(search FactSearch) (float64, error) {
	f.Taxonomy.Get(search.Taxonomy).Get(search.Field).
}
