package main

import (
	"fmt"
	"github.com/oliverasistemas/webcrawler/pagedownloader"
	"github.com/oliverasistemas/webcrawler/url_extractor"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// SafeVisited is a thread-safe data structure for storing visited URLs.
// It uses a map to store the URLs and a sync.Mutex to ensure concurrent access safety.
type SafeVisited struct {
	mu      sync.Mutex      // Mutex for synchronizing access to the visited map
	visited map[string]bool // Map for storing visited URLs
}

// Add method adds a URL to the SafeVisited data structure in a thread-safe manner.
func (sv *SafeVisited) Add(url string) {
	sv.mu.Lock()           // Lock the mutex to ensure exclusive access to the visited map
	defer sv.mu.Unlock()   // Unlock the mutex when the function returns
	sv.visited[url] = true // Mark the URL as visited
}

// Has method checks if a URL is already present in the SafeVisited data structure
// in a thread-safe manner. It returns true if the URL is present, false otherwise.
func (sv *SafeVisited) Has(url string) bool {
	sv.mu.Lock()           // Lock the mutex to ensure exclusive access to the visited map
	defer sv.mu.Unlock()   // Unlock the mutex when the function returns
	return sv.visited[url] // Return the value (true/false) associated with the URL in the visited map
}

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

func NewSafeVisited() *SafeVisited {
	return &SafeVisited{visited: make(map[string]bool)}
}

func main() {
	startingURL, destDir := parseArgs()
	safeVisited := NewSafeVisited()
	urlsChan := make(chan string)
	wg := setupCrawler(startingURL, destDir, safeVisited, urlsChan)
	handleInterrupt(wg, urlsChan)
	printCrawledURLs(urlsChan)
}

func parseArgs() (string, string) {
	if len(os.Args) != 3 {
		fmt.Println("Usage: webcrawler <starting-url> <destination-directory>")
		os.Exit(1)
	}
	return os.Args[1], os.Args[2]
}

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

// handleInterrupt sets up a signal handler to gracefully shut down the web crawler
// when an interrupt signal (Ctrl+C) or SIGTERM is received. It takes a sync.WaitGroup
// and a channel of strings (URLs) as input.
func handleInterrupt(wg *sync.WaitGroup, urlsChan chan string) {
	// Create a channel to listen for interrupt signals (e.g., Ctrl+C or SIGTERM).
	interruptChan := make(chan os.Signal, 1)

	// Register signal.Notify to listen for specific signals (os.Interrupt and syscall.SIGTERM)
	// and send them to the interruptChan.
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGTERM)

	// Start an anonymous goroutine to handle the interrupt signal.
	go func() {
		// Wait for an interrupt signal to be received on the interruptChan.
		<-interruptChan

		// Print a message indicating the crawler is stopping due to an interrupt signal.
		fmt.Println("\nReceived interrupt signal, stopping the crawler.")

		// Close the urlsChan to stop any ongoing crawling operations.
		close(urlsChan)

		// Wait for all crawling goroutines to finish.
		wg.Wait()

		// Exit the program with a status code of 0.
		os.Exit(0)
	}()
}

func printCrawledURLs(urlsChan chan string) {
	for url := range urlsChan {
		fmt.Println("Crawled URL:", url)
	}
}
