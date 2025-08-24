package xbrl

type XBRLCalcDoc struct {
	RoleRef []struct {
		URI  string `xml:"roleURI,attr"`
		Href string `xml:"href,attr"`
	} `xml:"roleRef"`
	CalculationLinks []struct {
		Loc []struct {
			Label string `xml:"xlink:label,attr"`
			Href  string `xml:"xlink:href,attr"`
		} `xml:"loc"`
		CalculationArc []struct {
			Order  string `xml:"order,attr"`
			Weight string `xml:"weight,attr"`
			From   string `xml:"xlink:from,attr"`
			To     string `xml:"xlink:to,attr"`
		} `xml:"calculationArc"`
	} `xml:"calculationLink"`
}
