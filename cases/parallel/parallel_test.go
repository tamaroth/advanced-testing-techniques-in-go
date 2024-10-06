package parallel

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func CanBeExecutedInParallel(source int) int {
	time.Sleep(time.Second)
	return source * 2
}

func TestCanBeExecutedInParallel(t *testing.T) {
	testCases := []struct {
		name   string
		source int
		expect int
	}{
		{
			name:   "test 1",
			source: 1,
			expect: 2,
		},
		{
			name:   "test 2",
			source: 2,
			expect: 4,
		},
		{
			name:   "test 3",
			source: 3,
			expect: 6,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			result := CanBeExecutedInParallel(testCase.source)
			assert.Equal(t, testCase.expect, result)
		})
	}
}

func TestCannotBeExecutedInParallel(t *testing.T) {
	testCases := []struct {
		name   string
		source int
		expect int
	}{
		{
			name:   "test 1",
			source: 1,
			expect: 2,
		},
		{
			name:   "test 2",
			source: 2,
			expect: 4,
		},
		{
			name:   "test 3",
			source: 3,
			expect: 6,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			result := CanBeExecutedInParallel(testCase.source)
			assert.Equal(t, testCase.expect, result)
		})
	}
}
