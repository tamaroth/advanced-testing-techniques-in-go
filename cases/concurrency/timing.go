package concurrency

import "sync"

type Aggregator struct {
	results []error
	mu      sync.Mutex
}

func (a *Aggregator) AddResult(err error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.results = append(a.results, err)
}

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
