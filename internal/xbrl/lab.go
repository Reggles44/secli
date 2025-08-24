package xbrl

type XBRLLabDoc struct {
	RoleRef []struct {
		URI string `xml:"roleURI,attr"`
	} `xml:"roleRef"`

	LabelLink []struct {
		ID    string `xml:"id,attr"`
		Label string `xml:"xlink:label,attr"`
		Role  string `xml:"xlink:role,attr"`
		Value string `xml:""`
	} `xml:"labelLink"`
}
