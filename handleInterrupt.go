package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

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
