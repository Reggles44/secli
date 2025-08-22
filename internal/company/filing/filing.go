package filing

import (
	"fmt"
	"strings"
	"time"

	"github.com/Reggles44/secli/internal/utils/request"
)

var (
	filingUrl           = "https://www.sec.gov/Archives/edgar/data/%v/%v/%v"
	filingCacheDuration = -1
)

type FilingMeta struct {
	CIK                   string
	AccessionNumber       string
	FilingDate            time.Time
	ReportDate            time.Time
	AcceptanceDateTime    time.Time
	Act                   string
	Form                  string
	FileNumber            string
	FilmNumber            string
	Items                 string
	Core_type             string
	Size                  int64
	IsXBRL                int8
	IsInlineXBRL          int8
	PrimaryDocument       string
	PrimaryDocDescription string
}

type FilingData struct {
	
	
}

func (f FilingMeta) GetFile() (*[]byte, error) {
	url := fmt.Sprintf(filingUrl, f.CIK, strings.Replace(f.AccessionNumber, "-", "", -1), f.PrimaryDocument)
	data, err := request.Get("GET", url, filingCacheDuration)
	if err != nil {
		return nil, err
	}

	return 
}
