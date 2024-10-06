package concurrency

import (
	"sync"
)

func worker(wg *sync.WaitGroup, perWorker int, counter *int) {
	defer wg.Done()
	for i := 0; i < perWorker; i++ {
		*counter++
	}
}

func DangerousRaceCondition(workerCount int, perWorker int) int {
	var wg sync.WaitGroup
	var counter int

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(&wg, perWorker, &counter)
	}

	wg.Wait()
	return counter
}
