package request

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"path"

	"github.com/Reggles44/secli/internal/cache"
	"github.com/spf13/viper"
)

func MakeUrl(tpl template.Template, data any) string {
	var buf bytes.Buffer
	tpl.Execute(&buf, data)
	return buf.String()
}

// TODO: Create requests per second limitation
func Get(method string, url string, cacheDuration int) (*[]byte, error) {
	// Read from cache
	data, err := cache.Read(cacheFilePath, cacheDuration)
	
	// If we could not read from cache
	if err != nil {

		// Fetch data
		data, err := execute(method, url)
		if err != nil {
			return nil, err
		}

		// Write to cache
		go cache.Write(cacheFileName, data)
	}

	return data, nil
}

func DownloadFile(method string, url string) (string, error) {
	data, err := cache.Read(cacheFilePath, -1)
}

func execute(method string, url string) (*[]byte, error) {
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
	if err != nil {
		return nil, err
	}

	return &data, nil
}
