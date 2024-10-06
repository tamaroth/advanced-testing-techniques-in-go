package pbt

import (
	"testing"
	"testing/quick"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"github.com/stretchr/testify/assert"
)

// TestRectangleProperties validates properties of the Rectangle struct using testing/quick.
func TestRectangleProperties(t *testing.T) {
	// Property 1: Area should always be non-negative
	areaNonNegative := func(width, height int) bool {
		// Skip invalid rectangles
		if width < 0 || height < 0 {
			return true
		}
		rect := Rectangle{Width: width, Height: height}
		return rect.Area() >= 0
	}

	// Property 2: Perimeter should be at least twice the largest dimension
	perimeterValid := func(width, height int) bool {
		// Skip invalid rectangles
		if width < 0 || height < 0 {
			return true
		}
		rect := Rectangle{Width: width, Height: height}
		perimeter := rect.Perimeter()
		return perimeter >= 2*width && perimeter >= 2*height
	}

	// Property 3: Resize by a factor should multiply area by factor squared
	resizeArea := func(width, height, scaleFactor int) bool {
		// Skip invalid values
		if width < 0 || height < 0 || scaleFactor < 0 {
			return true
		}
		rect := Rectangle{Width: width, Height: height}
		initialArea := rect.Area()
		rect.Resize(scaleFactor)
		expectedArea := initialArea * scaleFactor * scaleFactor
		return rect.Area() == expectedArea
	}

	// Run all properties
	assert.NoError(t, quick.Check(areaNonNegative, nil))
	assert.NoError(t, quick.Check(perimeterValid, nil))
	assert.NoError(t, quick.Check(resizeArea, nil))
}

func TestRectanglePropertiesWithGopter(t *testing.T) {
	properties := gopter.NewProperties(nil)

	// Property 1: Area should always be non-negative
	properties.Property("Area is non-negative", prop.ForAll(
		func(width, height int) bool {
			// Skip invalid rectangles
			if width < 0 || height < 0 {
				return true
			}
			rect := Rectangle{Width: width, Height: height}
			return rect.Area() >= 0
		},
		gen.IntRange(0, 1000), // Width generator
		gen.IntRange(0, 1000), // Height generator
	))

	// Property 2: Perimeter should be at least twice the largest dimension
	properties.Property("Perimeter is at least twice the largest dimension", prop.ForAll(
		func(width, height int) bool {
			if width < 0 || height < 0 {
				return true
			}
			rect := Rectangle{Width: width, Height: height}
			perimeter := rect.Perimeter()
			return perimeter >= 2*width && perimeter >= 2*height
		},
		gen.IntRange(0, 1000), // Width generator
		gen.IntRange(0, 1000), // Height generator
	))

	// Property 3: Resize by a factor should multiply area by factor squared
	properties.Property("Resize scales area by factor squared", prop.ForAll(
		func(width, height, scaleFactor int) bool {
			if width < 0 || height < 0 || scaleFactor < 0 {
				return true
			}
			rect := Rectangle{Width: width, Height: height}
			initialArea := rect.Area()
			rect.Resize(scaleFactor)
			expectedArea := initialArea * scaleFactor * scaleFactor
			return rect.Area() == expectedArea
		},
		gen.IntRange(0, 1000), // Width generator
		gen.IntRange(0, 1000), // Height generator
		gen.IntRange(1, 10),   // Scale factor generator
	))

	// Run the tests
	properties.TestingRun(t)
}
