package main

import "sync"

// setupCrawler initializes a web crawler with a starting URL, destination directory,
// a thread-safe visited URLs set, and a channel for sending found URLs. It returns a
// sync.WaitGroup for tracking the completion of the crawling process.
func setupCrawler(startingURL, destDir string, safeVisited *SafeVisited, urlsChan chan string) *sync.WaitGroup {
	// Create a new sync.WaitGroup for managing goroutines.
	var wg sync.WaitGroup

	// Increment the WaitGroup counter by 1 as we are about to launch a new goroutine.
	wg.Add(1)

	// Start the crawl function in a new goroutine.
	go crawl(startingURL, startingURL, destDir, safeVisited, &wg, urlsChan)

	// Start an anonymous goroutine to wait for the WaitGroup to finish,
	// then close the urlsChan channel.
	go func() {
		wg.Wait()
		close(urlsChan)
	}()

	// Return a pointer to the WaitGroup so the caller can wait for all
	// crawling goroutines to complete.
	return &wg
}
