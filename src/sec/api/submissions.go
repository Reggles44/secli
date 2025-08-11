package secapi

import "fmt"


const baseSubmissionURL string = "https://data.sec.gov/submissions/CIK%v.json"

func GetSubmissions(cik float64) error {
	fmt.Printf("HELP cik=%v\n", cik)
	resp, err := Request("GET", fmt.Sprintf(baseSubmissionURL, cik), nil)
	if err != nil {
		return err
	}

	fmt.Println(resp.Body)


	return nil
}
