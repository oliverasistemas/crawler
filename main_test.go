package main

import (
	"os"
	"path/filepath"
	"sync"
	"testing"
)

func TestCrawl(t *testing.T) {
	urlStr := "https://www.truvity.com/"
	destDir := "temp"
	safeVisited := NewSafeVisited()
	wg := &sync.WaitGroup{}
	urlsChan := make(chan string, 100)

	wg.Add(1)
	go crawl(urlStr, urlStr, destDir, safeVisited, wg, urlsChan)

	wg.Wait()
	close(urlsChan)

	// Check if the file with the specified path has been created.
	expectedFilePath := filepath.Join(destDir, "www.truvity.com", "blog", "ssi-digest-april-2022.html")
	if _, err := os.Stat(expectedFilePath); os.IsNotExist(err) {
		t.Errorf("File not found: %s", expectedFilePath)
	}
}
