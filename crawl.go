package main

import (
	"fmt"
	"github.com/oliverasistemas/webcrawler/pagedownloader"
	"github.com/oliverasistemas/webcrawler/url_extractor"
	"log"
	"sync"
)

// crawl is a function that recursively crawls a website starting from the given urlStr,
// downloading the content of each visited URL and saving it to the specified destDir.
// It ensures that URLs are not visited multiple times by using the safeVisited data structure.
// The function uses a sync.WaitGroup (wg) to keep track of running goroutines and a channel (urlsChan)
// to send the extracted URLs to another process.
//
// Parameters:
// - urlStr (string): The URL to start crawling from.
// - destDir (string): The directory where the downloaded content should be saved.
// - safeVisited (*SafeVisited): A thread-safe data structure to store visited URLs.
// - wg (*sync.WaitGroup): A wait group to track the number of active goroutines.
// - urlsChan (chan<- string): A channel to send the extracted URLs to another process.
func crawl(urlStr, baseURL, destDir string, safeVisited *SafeVisited, wg *sync.WaitGroup, urlsChan chan<- string) {
	// Ensure that the WaitGroup counter is decremented when the function returns.
	defer wg.Done()

	// Check if the URL has already been visited. If not, proceed with crawling.
	if !safeVisited.Has(urlStr) {
		fmt.Println("Crawling:", urlStr)
		safeVisited.Add(urlStr)

		// Download the content of the URL and save it to the destination directory.
		body, err := pagedownloader.DownloadPage(urlStr, destDir)
		if err != nil {
			log.Println("Error:", err)
			return
		}

		// Extract valid URLs from the downloaded content.
		urls, err := url_extractor.ExtractValidUrls(baseURL, body)
		if err != nil {
			log.Println("Error:", err)
			return
		}

		// Iterate through the extracted URLs and recursively crawl each one.
		// Add each URL to the urlsChan for further processing.
		for _, childUrl := range urls {
			wg.Add(1)
			go crawl(childUrl, baseURL, destDir, safeVisited, wg, urlsChan)
		}
	}
}
