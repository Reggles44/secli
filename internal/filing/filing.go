package filing

import (
	"fmt"
	"strings"
	"time"

	"github.com/Reggles44/secli/internal/cache"
	"github.com/Reggles44/secli/internal/xbrl"
)

var (
	filingUrl           = "https://www.sec.gov/Archives/edgar/data/%v/%v/%v-xbrl.zip"
	filingCacheDuration = -1
)

type Filing struct {
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
	Documents             map[string]*[]byte
}

func (f Filing) GetFilingDocuments() (map[string]*[]byte, error) {
	url := fmt.Sprintf(
		filingUrl,
		f.CIK,
		strings.ReplaceAll(f.AccessionNumber, "-", ""),
		f.AccessionNumber,
	)

	filingCache := cache.ZipCache{URL: url}
	filings, err := filingCache.Read()
	if err != nil {
		return nil, err
	}

	xbrl.XBRLRaw{}.Create(filings)

	return filings, nil
}
