package pagedownloader

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// downloadPage downloads the content of the specified URL and saves it to the
// destination directory. It returns the downloaded content as a string and
// any error encountered.
func DownloadPage(urlStr, destDir string) (string, error) {
	// Send an HTTP GET request to the specified URL
	resp, err := http.Get(urlStr)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body into a byte slice
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Parse the URL to obtain its components
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	// Create the destination file path based on the URL components
	fileName := filepath.Join(destDir, parsedURL.Host, parsedURL.Path)
	if strings.HasSuffix(fileName, "/") {
		fileName = filepath.Join(fileName, "index.html")
	} else {
		fileName = fileName + ".html"
	}

	// Create the necessary directories for the destination file path
	err = os.MkdirAll(filepath.Dir(fileName), 0755)
	if err != nil {
		return "", err
	}

	// Write the downloaded content to the destination file
	err = os.WriteFile(fileName, body, 0644)
	if err != nil {
		return "", err
	}

	// Return the downloaded content as a string
	return string(body), nil
}
