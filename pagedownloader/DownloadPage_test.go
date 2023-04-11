package pagedownloader

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestDownloadPage(t *testing.T) {
	// Create a test server that serves a simple HTML page
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("<html><body>Hello, World!</body></html>"))
		if err != nil {
			log.Println("error writing file")
		}
	}))
	defer ts.Close()

	// Define the destination directory for the downloaded page
	destDir := "test_downloads"

	// Call the DownloadPage function
	content, err := DownloadPage(ts.URL, destDir)
	if err != nil {
		t.Errorf("DownloadPage failed: %v", err)
	}

	// Check if the downloaded content matches the served content
	expectedContent := "<html><body>Hello, World!</body></html>"
	if content != expectedContent {
		t.Errorf("Downloaded content does not match expected content")
	}

	// Read the downloaded file and compare its content to the served content
	filePath := destDir + "/" + strings.TrimPrefix(ts.URL, "http://") + ".html"
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		t.Errorf("Failed to read downloaded file: %v", err)
	}

	if string(fileContent) != expectedContent {
		t.Errorf("Downloaded file content does not match expected content")
	}

	// Clean up the downloaded file and directory
	err = os.RemoveAll(destDir)
	if err != nil {
		t.Errorf("Failed to clean up test_downloads directory: %v", err)
	}
}
