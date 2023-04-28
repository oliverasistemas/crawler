package main

import (
	"fmt"
	"os"
)

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

func printCrawledURLs(urlsChan chan string) {
	for url := range urlsChan {
		fmt.Println("Crawled URL:", url)
	}
}
