package pbt_test

import (
	"reflect"
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"

	"github.com/tamaroth/advanced-testing-techniques-in-go/cases/pbt"
)

func TestRange(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	// Create a generator for pairs of integers where start <= end
	rangeGen := gen.IntRange(0, 100).FlatMap(
		func(start interface{}) gopter.Gen {
			s := start.(int)
			return gen.IntRange(s, 100).Map(
				func(end interface{}) []int {
					e, ok := end.(*gopter.GenResult).Retrieve()
					if !ok {
						return []int{}
					}
					return pbt.Range(start.(int), e.(int))
				})
		},
		reflect.TypeOf([]int{}),
	)

	properties.Property("Range should return a slice in the specified range",
		prop.ForAll(
			func(result []int) bool {
				if len(result) == 0 {
					return true // Empty slice is fine for invalid ranges
				}
				for i := 1; i < len(result); i++ {
					if result[i] != result[i-1]+1 {
						return false
					}
				}
				return true
			},
			rangeGen,
		))

	properties.TestingRun(t)
}
