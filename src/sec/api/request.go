package secapi

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

func Request(method string, url string, body io.Reader) (*http.Response, error) {
	// Create Request
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// TODO: Make company and email dynamic
	req.Header.Add("User-Agent", fmt.Sprintf("%s %s", "SEC CLI Tool", "reginaldbeakes@gmail.com"))

	// Debug Request
	reqDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("REQUEST:\n%s", string(reqDump))

	// Do Request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Debug Response
	respDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("RESPONSE:\n%s", string(respDump))

	return resp, nil
}
