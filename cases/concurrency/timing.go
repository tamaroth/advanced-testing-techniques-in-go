package concurrency

import "sync"

// Aggregator holds a slice to store results from goroutines
type Aggregator struct {
	results []error
	mu      sync.Mutex
}

// AddResult adds a result to the aggregator
func (a *Aggregator) AddResult(err error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.results = append(a.results, err)
}

// ResultsCount returns the number of results
func (a *Aggregator) FailedResults() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	var count int
	for _, err := range a.results {
		if err != nil {
			count++
		}
	}
	return count
}
