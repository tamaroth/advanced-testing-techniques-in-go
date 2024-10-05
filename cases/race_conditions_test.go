package cases_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tamaroth/advanced-testing-techniques-in-go/cases"
)

func TestDangerousRaceCondition(t *testing.T) {
	workers := 10
	perWorker := 1000

	counter := cases.DangerousRaceCondition(workers, perWorker)

	assert.Equal(t, workers*perWorker, counter)
}
