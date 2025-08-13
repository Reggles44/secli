package secapi

import "errors"

func (submission *Submission) GetFile(formSearch string) (*[]byte, error) {
	// for _, form := range submission.Filings.Recent.Form {
	// 	if form == formSearch {
	// 		// return Request("GET", )
	// 	}
	// }

	return nil, errors.New("file not found")
}
