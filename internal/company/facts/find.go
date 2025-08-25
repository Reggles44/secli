package facts

import (
	"slices"
	"time"
)

type FactFilter struct {
	Form   string
	Fields []string
}

func (f Facts) Find(form string, fields ...string) map[string]map[string]Value {
	found := make(map[string][]Value)

	for name, fact := range f.findFacts(fields...) {
		found[name] = fact.filterValues(form)
	}

	indexed := make(map[string]map[string]Value)
	for name, vs := range found {
		for _, v := range vs {
			if _, ok := indexed[v.ACCN]; !ok {
				indexed[v.ACCN] = make(map[string]Value)
			}
			indexed[v.ACCN][name] = v
		}
	}

	return indexed
}

type foundFact struct {
	Name string
	Fact Fact
}

func (f Facts) findFacts(fields ...string) map[string]Fact {
	found := make(map[string]Fact)

	for _, fmap := range f.Data {
		for f, fc := range fmap {
			if slices.Contains(fields, f) {
				found[f] = fc
			}
		}
	}

	return found
}

func (f Fact) filterValues(form string) []Value {
	var values []Value

	for _, value := range f.Units["USD"] {
		if value.Form == form {
			if value.Start != "" && value.End != "" {
				start, _ := time.Parse(time.DateOnly, value.Start)
				end, _ := time.Parse(time.DateOnly, value.End)
				dif := end.Sub(start)
				mdif := int64(dif.Hours() / 24 / 30)
				if (value.Form == "10-Q" && mdif == 3) || (form == "10-K" && mdif == 12) {
					values = append(values, value)
				}
			} else {
				values = append(values, value)
			}
		}
	}

	return values
}
