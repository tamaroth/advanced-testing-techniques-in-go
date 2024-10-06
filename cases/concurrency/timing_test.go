package concurrency

import (
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAggregatorWithoutWaiting(t *testing.T) {
	aggregator := &Aggregator{}
	numGoroutines := 100

	for i := 0; i < numGoroutines; i++ {
		go func() {
			val := errors.New("test error")
			time.Sleep(time.Second) // Simulate work
			aggregator.AddResult(val)
		}()
	}

	assert.Equal(t, 0, aggregator.FailedResults())
}

func TestAggregatorWithWaiting(t *testing.T) {
	aggregator := &Aggregator{}
	numGoroutines := 100
	wg := sync.WaitGroup{}

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			val := errors.New("test error")
			time.Sleep(time.Second) // Simulate work
			aggregator.AddResult(val)
		}()
	}

	wg.Wait()
	assert.Equal(t, 0, aggregator.FailedResults())
}
