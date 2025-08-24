package xbrl

type XBRLDefDoc struct {
	ArcRoleRef []struct {
		URI string `xml:"arcoroleURI,attr"`
	} `xml:"arcroleRef"`

	DefinitionLink []struct {
		Loc []struct {
			Label string `xml:"xlink:label,attr"`
			Href  string `xml:"xlink:href,attr"`
		} `xml:"loc"`

		DefinitionArc []struct {
			ArcRole string `xml:"xlink:arcrole,attr"`
			From    string `xml:"xlink:from,attr"`
			To      string `xml:"xlink:to,attr"`
			Order   string `xml:"order,attr"`
		} `xml:"definitionArc"`
	} `xml:"definitionLink"`
}
