package xbrl

type XBRLSchemaDoc struct {
	Imports []struct {
		Namespace      string `xml:"namespace,attr"`
		SchemaLocation string `xml:"schemaLocation,attr"`
	} `xml:"import"`

	Annocations struct {
		AppInfo struct {
			RoleTypes []struct {
				ID         string   `xml:"id,attr"`
				Definition string   `xml:"definition"`
				UsedOn     []string `xml:"usedOn"`
			} `xml:"roleType"`
		} `xml:"appinfo"`
	} `xml:"annotation"`

	Elements []struct {
		ID   string `xml:"id,attr"`
		Name string `xml:"name,attr"`
	} `xml:"element"`
}
