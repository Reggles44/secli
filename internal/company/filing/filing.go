package filing

type Filing struct {
	AccessionNumber       string `json:"accessionNumber"`
	FilingDate            string `json:"filingDate"`
	ReportDate            string `json:"reportDate"`
	AcceptanceDateTime    string `json:"acceptanceDateTime"`
	Act                   string `json:"act"`
	Form                  string `json:"form"`
	FileNumber            string `json:"fileNumber"`
	FilmNumber            string `json:"filmNumber"`
	Items                 string `json:"items"`
	Core_type             string `json:"core_type"`
	Size                  int64  `json:"size"`
	IsXBRL                int8   `json:"isXBRL"`
	IsInlineXBRL          int8   `json:"isInlineXBRL"`
	PrimaryDocument       string `json:"primaryDocument"`
	PrimaryDocDescription string `json:"primaryDocDescription"`
}
