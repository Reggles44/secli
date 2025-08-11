package secapi

import "fmt"


const baseSubmissionURL string = "https://data.sec.gov/submissions/CIK%010d.json"

func GetSubmissions(cik int) error {
	fmt.Printf("HELP cik=%v\n", cik)

	url := fmt.Sprintf(baseSubmissionURL, cik)
	fmt.Printf("HELP url=%v\n", url)

	resp, err := Request("GET", url, nil)
	if err != nil {
		return err
	}

	fmt.Println(resp.Body)


	return nil
}
