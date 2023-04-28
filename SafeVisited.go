package main

import "sync"

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

func NewSafeVisited() *SafeVisited {
	return &SafeVisited{visited: make(map[string]bool)}
}
