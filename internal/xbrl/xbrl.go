package xbrl

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"regexp"

	"golang.org/x/net/html/charset"
)

var schemaRegex = regexp.MustCompile(`([a-z]+)-(\d{8})\.xsd`)

// calcRegex   = regexp.MustCompile("[a-z]+_d{8}_calc.xml")
// defRegex    = regexp.MustCompile("[a-z]+_d{8}_def.xml")
// labRegex    = regexp.MustCompile("[a-z]+_d{8}_lab.xml")
// preRegex    = regexp.MustCompile("[a-z]+_d{8}_pre.xml")
// linkRegex   = regexp.MustCompile("[a-z]+_d{8}.xml")

type XBRLRaw struct {
	Schema XBRLSchemaDoc
	// Calc   XBRLCalcDoc
	// Def    XBRLDefDoc
	// Lab    XBRLLabDoc
	// Pre    XBRLPreDoc
	// Link   XBRLLinkDock
}

func matchUnmarshal(re *regexp.Regexp, name string, data *[]byte, v any) {
	if re.MatchString(name) {
		fmt.Println(name)
		reader := bytes.NewReader(*data)
		decoder := xml.NewDecoder(reader)
		decoder.CharsetReader = charset.NewReaderLabel
		err := decoder.Decode(v)
		if err != nil {
			panic(err)
		}
	}
}

func (x XBRLRaw) Create(files map[string]*[]byte) XBRLRaw {
	for fn, d := range files {
		matchUnmarshal(schemaRegex, fn, d, &x.Schema)
		// matchUnmarshal(calcRegex, fn, d, &x.Calc)
		// matchUnmarshal(defRegex, fn, d, &x.Def)
		// matchUnmarshal(labRegex, fn, d, &x.Lab)
		// matchUnmarshal(preRegex, fn, d, &x.Pre)
		// matchUnmarshal(linkRegex, fn, d, &x.Link)
	}

	fmt.Println(x.Schema.Annocations.AppInfo.RoleTypes[0])
	fmt.Println(x.Schema.Elements[0].ID)

	return x
}
