package concurrency

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDangerousRaceCondition(t *testing.T) {
	workers := 10
	perWorker := 1000

	counter := DangerousRaceCondition(workers, perWorker)

	assert.Equal(t, workers*perWorker, counter)
}
