package request

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/spf13/viper"
)

func Do(method string, url string) ([]byte, error) {
	// Create Request
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	// TODO: Make company and email dynamic
	req.Header.Add("User-Agent", fmt.Sprintf("%s %s", "SEC CLI Tool", "reginaldbeakes@gmail.com"))

	// Debug Request
	if viper.GetBool("debug") {
		reqDump, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("REQUEST:\n%s", string(reqDump))
	}

	// Do Request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Debug Response
	if viper.GetBool("debug") {
		respDump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("RESPONSE:\n%s", string(respDump))
	}

	data, err := io.ReadAll(resp.Body)
	return data, err
}
